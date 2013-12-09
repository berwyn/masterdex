package middleware

import (
	html "html/template"
	"log"
	"net/http"
	. "reflect"
	"regexp"
	"strings"
)

type ResponseTypeHandler struct {
	controller     interface{}
	methodMap      map[string]Method
	descriptionMap map[string]string
	templateName   string
	template       *html.Template
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
			h.methodMap[parts[0]] = t.MethodByName(strings.Title(field.Name))
			h.descriptionMap[parts[0]] = parts[1]
		}
	}
	return &h
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
		}

		return
	}
}
