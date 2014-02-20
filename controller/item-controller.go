package controller

import (
	"github.com/codegangsta/martini"
)

type ItemController struct {
	datastore Datastore
}

func (ctrl ItemController) Register(server *martini.ClassicMartini) {
	server.Get("/item", ctrl.Index)
}

func (ctrl ItemController) Index(request *Request) {
	request.Data = new(struct{})
	request.Status = 200
	request.Template = "item"
}
