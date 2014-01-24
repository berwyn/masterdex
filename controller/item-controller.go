package controller

import (
	"github.com/codegangsta/martini"
	"github.com/eaigner/hood"
)

type ItemController struct {
	Database *hood.Hood
}

func (ctrl ItemController) Register(server *martini.ClassicMartini) {
	server.Get("/item", ctrl.Index)
}

func (ctrl ItemController) Index(request *Request) {
	request.Data = new(struct{})
	request.Status = 200
	request.Template = "item"
}
