// Package controller defines all our Martini controllers
package controller

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	mime_type_html = "text/html"
	mime_type_json = "application/json"
)

var (
	html_regex = regexp.MustCompile(`text\/html`)
	json_regex = regexp.MustCompile(`application\/json`)
)

// Definition contract for all controllers
type Controller interface {
	Register(*martini.ClassicMartini)
}

// Definition contract for all datastores. Specifically,
// this was implemented so that testing was much more feasible,
// As well as decoupling controllers from datastores making
// possible future migrations easier
type Datastore interface {
	// Given a unique string ID, returns the first entity to match
	Find(id string) (interface{}, error)
	// Given an entity, this will insert new entries or update existing entries
	Insert(entity interface{}) (interface{}, error)
	// Given an entity, this will update existing ones. Use this if you want
	// an explicit failure for non-existing entities
	Update(entity interface{}) (interface{}, error)
	// Given a unique string ID, this deletes the entity in the
	// datastore
	Delete(id string) error
}

// This is our wrapper for the transaction
type Request struct {
	// This is the payload to return to the client
	Data interface{}
	// This is the status to write back
	Status int
	// In the case of an HTML request, this is the template
	// To use
	Template string
	// This is so that controller know whether to use a JSON implementation
	UsingJSON bool
	// Whether the request contains a JSON payload
	ContainsJSON bool
	// Any client-sent data
	Payload interface{}
	// MimeType of response to render
	ResponseType string
}

// Convenience method to set an error status, as well as provide a JSON error body
func (req *Request) Error(status int, message string) {
	req.Status = status
	req.Data = map[string]string{
		"error": message,
		"code":  strconv.Itoa(status),
	}
}

// This is a Martini middle-ware to wrap requests, provide our request wrapper
// to controllers, as well as serialise the response. This removes a tonne of
// boilerplate from the controller method bodies
func JsonRequstRouter(c martini.Context, request *http.Request, r render.Render) {
	body := &Request{
		Status: 0,
	}
	setResponseType(body, request)
	setPayload(body, request)
	if body.Status != 0 {
		return
	}

	c.Map(body)
	c.Next()

	if body.Data == nil {
		body.Data = new(struct{})
	}
	switch body.ResponseType {
	default:
		fallthrough
	case mime_type_html:
		r.HTML(body.Status, body.Template, body.Data)
	case mime_type_json:
		r.JSON(body.Status, body.Data)
	}
}

func setResponseType(body *Request, request *http.Request) {
	for _, part := range strings.Split(request.Header.Get("Accept"), ",") {
		fmt.Println(part)
		switch {
		default:
			fallthrough
		case html_regex.MatchString(part):
			body.ResponseType = mime_type_html
			break
		case json_regex.MatchString(part):
			body.ResponseType = mime_type_json
			break
		}
	}
}

func setPayload(body *Request, request *http.Request) {
	if request.Header.Get("Content-Type") == "application/json" {
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
			body.Error(http.StatusInternalServerError, "We could not process your request, please try again later")
			return
		}
		body.ContainsJSON = true
		body.Payload = payload
	}
	// We can add more options here later, for example yaml
}
