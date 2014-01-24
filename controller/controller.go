package controller

import (
	"github.com/codegangsta/martini"
	"net/http"
)

type Controller interface {
	Register(*martini.ClassicMartini)
}

type Request struct {
	Data     interface{}
	Status   int
	Template string `json:"-"`
}

func HasJSON(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/json"
}

func UseJSON(req *http.Request) bool {
	return req.Header.Get("Accept") == "application/json"
}
