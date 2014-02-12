// Package controller defines all our Martini controllers
package controller

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"net/http"
	"strconv"
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
	Find(id string) (interface{}, error)
	Insert(entity interface{}) (interface{}, error)
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
		UsingJSON:    useJSON(request),
		ContainsJSON: hasJSON(request),
	}
	if body.ContainsJSON {
		// TODO Come back and error handle here
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
		}
		body.Payload = payload
	}
	c.Map(body)

	c.Next()

	if body.Data == nil {
		body.Data = new(struct{})
	}
	if body.UsingJSON {
		r.JSON(body.Status, body.Data)
	} else {
		r.HTML(body.Status, body.Template, body.Data)
	}
}

// Convenience method to check a request's Content-Type header and determine
// whether it contains a JSON payload
func hasJSON(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/json"
}

// Convenience method to check a request's Accept header and
// determine if we should send a JSON payload
func useJSON(req *http.Request) bool {
	return req.Header.Get("Accept") == "application/json"
}
