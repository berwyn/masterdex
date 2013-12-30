package controller

import (
	"github.com/codegangsta/martini"
	"net/http"
)

type Controller interface {
	Register(*martini.ClassicMartini)
}

func hasJson(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/json"
}

func useJson(req *http.Request) bool {
	return req.Header.Get("Accept") == "application/json"
}
