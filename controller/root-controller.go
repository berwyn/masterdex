package controller

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

type RootController struct{}

func (ctrl RootController) Register(server *martini.ClassicMartini) {
	server.Get("/", func(r render.Render) {
		r.HTML(200, "root", map[string]interface{}{"cssClass": "home"})
	})
}
