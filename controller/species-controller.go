package controller

import (
	. "../model"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/eaigner/hood"
	"io/ioutil"
	"log"
	"net/http"
)

type SpeciesController struct {
	Database *hood.Hood
}

func (ctrl SpeciesController) Register(server *martini.ClassicMartini) {
	server.Get("/pokemon/:dex/:id", ctrl.Read)
	server.Post("/pokemon", ctrl.Create)
	server.Put("/pokemon/:dex/:id", ctrl.Update)
	server.Patch("/pokemon/:dex/:id", ctrl.Update)
	server.Delete("/pokemon/:dex/:id", ctrl.Delete)
	server.Options("/pokemon", ctrl.Metadata)
}

func (ctrl SpeciesController) Create(r render.Render, req *http.Request, logger *log.Logger) {
	logger.Println("Got to POST /pokemon")
	hasJson := req.Header.Get("Content-Type") == "application/json"
	if hasJson {
		var body map[string]string
		data, readErr := ioutil.ReadAll(req.Body)
		jsonErr := json.Unmarshal(data, &body)
		fmt.Println(fmt.Sprintf("Body: %s\nError: %s\nRead error: %s", body, jsonErr, readErr))
		r.JSON(http.StatusCreated, new(interface{}))
	} else {
		req.ParseForm()
		fmt.Println(fmt.Sprintf("Form Body: %s", req.Form))
		r.HTML(http.StatusCreated, "species", new(interface{}))
	}
}

func (ctrl SpeciesController) Read(params martini.Params, r render.Render, req *http.Request) {
	useJson := req.Header.Get("Accept") == "application/json"
	var results []Species
	if params["dex"] == "national" {
		err := ctrl.Database.Where("dex_number", "=", params["id"]).Limit(1).Find(&results)
		if err == nil {
			if useJson {
				r.JSON(http.StatusOK, results[0])
			} else {
				r.HTML(http.StatusOK, "species", results[0])
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

func (ctrl SpeciesController) Metadata(r render.Render) {
	options := make(map[string]string)
	options["test"] = "test"
	r.JSON(http.StatusOK, options)
}
