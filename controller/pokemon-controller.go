package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/berwyn/masterdex/model"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/eaigner/hood"
	"io/ioutil"
	"log"
	"net/http"
)

type PokemonController struct {
	Database *hood.Hood
}

func hasJson(req *http.Request) bool {
	return req.Header.Get("Content-Type") == "application/json"
}

func useJson(req *http.Request) bool {
	return req.Header.Get("Accept") == "application/json"
}

func (ctrl PokemonController) Register(server *martini.ClassicMartini) {
	server.Get("/pokemon", ctrl.Index)
	server.Get("/pokemon/:dex/:id", ctrl.Read)
	server.Post("/pokemon", ctrl.Create)
	server.Put("/pokemon/:dex/:id", ctrl.Update)
	server.Patch("/pokemon/:dex/:id", ctrl.Update)
	server.Delete("/pokemon/:dex/:id", ctrl.Delete)
	server.Options("/pokemon", ctrl.Metadata)
}

func (ctrl PokemonController) Index(r render.Render, req *http.Request) {
	if useJson(req) {
		r.Error(http.StatusTeapot)
	} else {
		r.HTML(200, "pokemon", nil)
	}
}

func (ctrl PokemonController) Create(r render.Render, w http.ResponseWriter, req *http.Request, logger *log.Logger) {
	var payload Species
	if hasJson(req) {
		data, readErr := ioutil.ReadAll(req.Body)
		jsonErr := json.Unmarshal(data, &payload)
		if jsonErr != nil || readErr != nil {
			r.Error(http.StatusInternalServerError)
			return
		}
	} else {
		r.Error(http.StatusTeapot)
	}

	if ctrl.exists("national", payload.DexNumber) {
		r.Error(http.StatusConflict)
		return
	}

	tx := ctrl.Database.Begin()
	tx.Save(&payload)
	err := tx.Commit()

	if err != nil {
		tx.Rollback()
		r.Error(422)
		return
	}

	if useJson(req) {
		r.JSON(http.StatusCreated, payload)
	} else {
		http.Redirect(w, req, fmt.Sprintf("/pokemon/national/%d", payload.DexNumber), http.StatusMovedPermanently)
	}
}

func (ctrl PokemonController) Read(params martini.Params, r render.Render, req *http.Request) {
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

func (ctrl PokemonController) Update(params martini.Params) {

}

func (ctrl PokemonController) Delete() {

}

func (ctrl PokemonController) Metadata(r render.Render) {
	options := make(map[string]string)
	options["test"] = "test"
	r.JSON(http.StatusOK, options)
}

func (ctrl PokemonController) exists(dex string, id int) bool {
	var result []Species
	if dex == "national" {
		ctrl.Database.Where("dex_number", "=", id).Limit(1).Find(&result)
		return len(result) > 0
	} else {
		return false
	}
}

func (ctrl PokemonController) loadOne(dex string, id int) (*Species, error) {
	var result []Species
	if dex == "national" {
		err := ctrl.Database.Where("dex_number", "=", id).Limit(1).Find(&result)
		if err != nil {
			return &result[0], nil
		}
		return nil, err
	}
	return nil, errors.New("NYI")
}
