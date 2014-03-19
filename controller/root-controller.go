package controller

import (
	"github.com/codegangsta/martini"
	"net/http"
)

type RootController struct{}

func (ctrl RootController) Register(server *martini.ClassicMartini) {
	server.Get("/", func(request *Request) {
		request.Status = http.StatusOK
		request.Template = "root"
	})
}
