package controller

import (
	. "../model"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/eaigner/hood"
	"net/http"
)

type SpeciesController struct {
	Database *hood.Hood
}

func (ctrl SpeciesController) Register(server *martini.ClassicMartini) {
	server.Get("/pokemon/:dex/:id", ctrl.Read)
	server.Post("/pokemon/", ctrl.Create)
	server.Patch("/pokemon/:dex/:id", ctrl.Update)
	server.Delete("/pokemon/:dex/:id", ctrl.Delete)
}

func (ctrl SpeciesController) Create() {

}

func (ctrl SpeciesController) Read(params martini.Params, r render.Render, req *http.Request) {
	useJson := req.Header.Get("Accept") == "application/json"
	var results []Species
	if params["dex"] == "national" {
		err := ctrl.Database.Where("dex_number", "=", params["id"]).Limit(1).Find(&results)
		if err == nil {
			if useJson {
				r.JSON(200, results[0])
			} else {
				r.HTML(200, "species", results[0])
			}
		} else {
			r.Error(500)
		}
	}
}

func (ctrl SpeciesController) Update(params martini.Params) {

}

func (ctrl SpeciesController) Delete() {

}
