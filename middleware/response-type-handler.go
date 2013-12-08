package middleware

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
)

type ResponseTypeHandler struct {
	loader       func(string) (interface{}, ResponseTypeHandlerError)
	templateName string
	template     *Template
}

type ResponseTypeHandlerError struct {
	errorCode    int
	errorMessage string
}

func NewResponseTypeHandler(tmpl string) *ResponseTypeHandler {
	return &ResponseTypeHandler{templateName: tmpl, template: template.Must(template.ParseFiles("../views/"+tmpl+".html", "../views/"+tmpl+".json"))}
}

func (h *ResponseTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	entity, err := h.loader("")
	accept := r.Header.Get("Accept")
	useJson, _ := regexp.MatchString("application/json|text/json", accept)

	if err != nil {
		w.Header(err.errorCode)
		if useJson {
			fmt.Fprintf(w, "{error_code:%d, error_message:%s}", err.errorCode, err.errorMessage)
		} else {
			http.Redirect(w, r, "/static/"+err.errorCode+".http", err.errorCode)
		}
	}

	if useJson {
		h.template.Execute(w, "../views/"+h.templateName+".json", entity)
	} else {
		h.template.Execute(w, "../views/"+h.templateName+".html", entity)
	}
}
