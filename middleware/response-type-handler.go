package middleware

import (
	"fmt"
	html "html/template"
	"net/http"
	. "reflect"
	"regexp"
	"strings"
)

type ResponseTypeHandler struct {
	controller   interface{}
	methodMap    map[string]Method
	templateName string
	template     *html.Template
}

type ResponseTypeHandlerError struct {
	ErrorCode    int
	ErrorMessage string
}

func NewResponseTypeHandler(controller interface{}, tmpl string) *ResponseTypeHandler {
	h := ResponseTypeHandler{
		controller:   controller,
		templateName: tmpl,
		template:     html.Must(html.ParseFiles("views/" + tmpl + ".html")),
	}
	t := TypeOf(controller)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("restr")
		if tag != "" {
			parts := strings.Split(tag, ",")
			method, found := t.MethodByName(strings.Title(field.Name))
			if !found {
				panic(fmt.Sprintf("Couldn't find the method %s on the type %s", strings.Title(field.Name), t.Name()))
			}
			h.methodMap[parts[0]] = method
		}
	}
	return &h
}

func (h *ResponseTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
