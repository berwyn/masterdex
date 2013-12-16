package middleware

import (
	"encoding/json"
	"fmt"
	html "html/template"
	"log"
	"net/http"
	"path"
	. "reflect"
	"regexp"
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
	error
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
	log.Println(fmt.Sprintf("[DEBUG] %s - :%d @%d", t.Name(), t.NumField(), t.NumMethod()))
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if tag != "" {
			methodName := strings.Title(field.Name)
			method, found := t.MethodByName(methodName)
			if !found {
				for j := 0; j < t.NumMethod(); j++ {
					log.Println(fmt.Sprintf("[ERROR] %s#%s", t.Name(), t.Method(j).Name))
				}
				panic(fmt.Sprintf("[ERROR] Couldn't find the method %s on the type %s", methodName, t.Name()))
			}
			log.Println(fmt.Sprintf("[DEBUG] \t#%s -> %s", methodName, tag))
			handler.methodMap[tag] = method
		}
	}
}

func (handler ResponseTypeHandler) checkAccept(acceptString string) bool {
	matched, _ := regexp.MatchString("application/json|text/json", acceptString)
	return matched
}

func (handler ResponseTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("%s - %s", r.Method, r.URL.Path))
	if _, exists := handler.methodMap[r.Method]; exists {
		args := []Value{ValueOf(handler.controller), ValueOf(path.Base(r.URL.Path))}
		values := handler.methodMap[r.Method].Func.Call(args)
		for i := 0; i < len(values); i++ {
			log.Println(fmt.Sprintf("%d: %s - %s", i, values[i], values[i].Interface()))
		}
		useJson := handler.checkAccept(r.Header.Get("Accept"))
		if useJson {
			body, err := json.Marshal(values[0].Interface())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, `{"error":"Internal server error"}`)
				return
			}
			fmt.Fprintf(w, "%s", body)
		} else {

		}
	} else {
		log.Println(fmt.Sprintf("%s not found in %s", r.Method, handler.methodMap))
		// TODO Handle unsupported methods
	}
}
