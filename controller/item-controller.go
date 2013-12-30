package controller

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/eaigner/hood"
	"net/http"
)

type ItemController struct {
	Database *hood.Hood
}

func (ctrl ItemController) Register(server *martini.ClassicMartini) {
	server.Get("/item", ctrl.Index)
}

func (ctrl ItemController) Index(r render.Render, req *http.Request) {
	if useJson(req) {
		r.Error(http.StatusTeapot)
	} else {
		r.HTML(http.StatusOK, "item", nil)
	}
}
