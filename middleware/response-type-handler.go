package middleware

import (
	"fmt"
	html "html/template"
	"log"
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
		methodMap:    make(map[string]Method),
	}
	h.mapFieldsToMethods()
	return &h
}

func (handler ResponseTypeHandler) mapFieldsToMethods() {
	t := TypeOf(handler.controller)
	log.Println(fmt.Sprintf("%s - :%d @%d", t.Name(), t.NumField(), t.NumMethod()))
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		log.Println(fmt.Sprintf("Processing %s:%s with tag %s", t.Name(), field.Name, tag))
		if tag != "" {
			methodName := strings.Title(field.Name)
			method, found := t.MethodByName(methodName)
			if !found {
				for j := 0; j < t.NumMethod(); j++ {
					log.Println(fmt.Sprintf("%s#%s", t.Name(), t.Method(j).Name))
				}
				panic(fmt.Sprintf("Couldn't find the method %s on the type %s", methodName, t.Name()))
			}
			log.Println(fmt.Sprintf("Mapping %s#%s to %s", t.Name(), methodName, tag))
			handler.methodMap[tag] = method
		}
	}
}

func (handler ResponseTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL.Path))
	if _, exists := handler.methodMap[r.Method]; exists {
		returns := handler.methodMap[r.Method].Func.Call([]Value{handler.controller, r.URL.Path})
		log.Println(fmt.Sprintf("%s", returns))
	} else {
		log.Println(fmt.Sprintf("%s not found in %s", r.Method, handler.methodMap))
		// TODO Handle unsupported methods
	}
}
