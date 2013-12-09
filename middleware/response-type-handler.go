package middleware

import (
	// "encoding/json"
	"fmt"
	html "html/template"
	// "log"
	"net/http"
	// "path"
	"regexp"
)

type ResponseTypeHandler struct {
	create       func(string) *ResponseTypeHandlerError
	read         func(string) (interface{}, *ResponseTypeHandlerError)
	update       func(string) *ResponseTypeHandlerError
	del          func(string) *ResponseTypeHandlerError
	descriptions map[string]string
	templateName string
	template     *html.Template
}

type ResponseTypeHandlerError struct {
	ErrorCode    int
	ErrorMessage string
}

func NewResponseTypeHandler(
	create func(string) *ResponseTypeHandlerError,
	read func(string) (interface{}, *ResponseTypeHandlerError),
	update func(string) *ResponseTypeHandlerError,
	del func(string) *ResponseTypeHandlerError,
	tmpl string,
	descriptor map[string]string,
) *ResponseTypeHandler {
	return &ResponseTypeHandler{
		read:         read,
		templateName: tmpl,
		template:     html.Must(html.ParseFiles("views/" + tmpl + ".html")),
	}
}

func (h *ResponseTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// id := path.Base(r.URL.Path)
	accept := r.Header.Get("Accept")
	useJson, _ := regexp.MatchString("application/json|text/json", accept)

	if match, _ := regexp.MatchString("GET|POST|PUT|PATCH|DELETE", r.Method); match {
		// entity, err := h.read(id)

		return
	}

	if match, _ := regexp.MatchString("OPTIONS", r.Method); match {
		if !useJson {
			http.Redirect(w, r, r.URL.Path, http.StatusTemporaryRedirect)
		} else {
			body := "{"
			if h.descriptions != nil {
				for k, v := range h.descriptions {
					body += fmt.Sprintf("\"%s\":{\"description\":\"%s\"}", k, v)
				}
			}
			body += "}"
			fmt.Fprintf(w, "%s", body)
		}

		return
	}

	// if useJson {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	log.Println("[DEBUG] Serving JSON for template " + h.templateName)
	// } else {
	// 	w.Header().Add("Content-Type", "text/html")
	// 	log.Println("[DEBUG] Serving HTML for template " + h.templateName)
	// }

	// if err != nil {
	// 	w.WriteHeader(err.ErrorCode)
	// 	if useJson {
	// 		fmt.Fprintf(w, "{error_code:%d, error_message:%s}", err.ErrorCode, err.ErrorMessage)
	// 	} else {
	// 		http.Redirect(w, r, fmt.Sprintf("/static/%d.http", err.ErrorCode), err.ErrorCode)
	// 	}
	// 	return
	// }

	// if useJson {
	// 	b, err := json.Marshal(entity)
	// 	if err == nil {
	// 		fmt.Fprintf(w, "%s", b)
	// 	} else {
	// 		fmt.Fprintf(w, "{error_code:%s, error_message:%s}", http.StatusInternalServerError, "Something went wrong")
	// 	}
	// } else {
	// 	h.template.Execute(w, entity)
	// }
}
