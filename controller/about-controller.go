package controller

import (
	"github.com/codegangsta/martini"
	"github.com/eaigner/hood"
)

type AboutController struct {
	Database *hood.Hood
}

func (ctrl AboutController) Register(server *martini.ClassicMartini) {
	server.Get("/about", ctrl.Index)
}

func (ctrl AboutController) Index(request *Request) {
	request.Data = new(struct{})
	request.Status = 200
	request.Template = "about"
}
