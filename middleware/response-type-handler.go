package middleware

import (
	"fmt"
	html "html/template"
	"net/http"
	. "reflect"
	"strings"
)

const tagName = "rest"

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
	h.mapFieldsToMethods()
	return &h
}

func (handler *ResponseTypeHandler) mapFieldsToMethods() {
	t := TypeOf(handler.controller)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("restr")
		if tag != "" {
			methodName := strings.Title(field.Name)
			method, found := t.MethodByName(methodName)
			if !found {
				panic(fmt.Sprintf("Couldn't find the method %s on the type %s", methodName, t.Name()))
			}
			handler.methodMap[tag] = method
		}
	}
}

func (handler *ResponseTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
