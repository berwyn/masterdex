package controller

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/eaigner/hood"
	"net/http"
)

type AboutController struct {
	Database *hood.Hood
}

func (ctrl AboutController) Register(server *martini.ClassicMartini) {
	server.Get("/about", ctrl.Index)
}

func (ctrl AboutController) Index(r render.Render, req *http.Request) {
	if useJSON(req) {
		r.Error(http.StatusTeapot)
	} else {
		r.HTML(http.StatusOK, "about", nil)
	}
}
