package controller

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Controller interface {
	Register(*martini.ClassicMartini)
}

type Datastore interface {
	Find(id string) (interface{}, error)
}

type Request struct {
	Data         interface{}
	Status       int
	Template     string
	UsingJSON    bool
	ContainsJSON bool
	Payload      interface{}
}

func (req *Request) Error(status int, message string) {
	req.Status = status
	req.Data = map[string]string{
		"error": message,
		"code":  strconv.Itoa(status),
	}
}

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

func hasJSON(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/json"
}

func useJSON(req *http.Request) bool {
	return req.Header.Get("Accept") == "application/json"
}
