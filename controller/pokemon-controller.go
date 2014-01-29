package controller

import (
	"encoding/json"
	"fmt"
	. "github.com/berwyn/masterdex/model"
	"github.com/codegangsta/martini"
	"github.com/eaigner/hood"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	ERROR_BAD_ID           = -1
	ERROR_BAD_REGION       = -2
	ERROR_ID_NOT_IN_REGION = -3
)

type PokemonController struct {
	Database *hood.Hood
}

type PokemonNotFoundError struct {
	id int
}

func (err PokemonNotFoundError) Error() string {
	return fmt.Sprintf("Could not find the pokemon with ID %d", err.id)
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

func (ctrl PokemonController) Index(response *Request) {
	if response.UsingJSON {
		ctrl.Metadata(response)
	} else {
		response.Status = 200
		response.Data = new(struct{})
		response.Template = "pokemon"
	}
}

func (ctrl PokemonController) Create(response *Request, logger *log.Logger) {
	if response.ContainsJSON {
		var pkmn Species
		var bytes = response.Payload.([]byte)
		jsonErr := json.Unmarshal(bytes, &pkmn)
		if jsonErr != nil {
			response.Error(422, "There was an issue with your JSON")
			return
		}

		tx := ctrl.Database.Begin()
		_, saveErr := tx.Save(&pkmn)
		commitErr := tx.Commit()
		if saveErr != nil || commitErr != nil {
			response.Error(http.StatusInternalServerError, "There was a problem with your request, please try again later")
			return
		}

		response.Status = http.StatusOK
	}
}

func (ctrl PokemonController) Read(params martini.Params, response *Request) {
	id := regionalIDToNational(params["region"], params["id"])
	switch id {
	case ERROR_BAD_ID:
		response.Error(http.StatusBadRequest, fmt.Sprintf(`'%s' is not a valid ID for a pokemon`, params["id"]))
	case ERROR_BAD_REGION:
		response.Error(http.StatusBadRequest, fmt.Sprintf(`'%s' is not a valid region`, params["region"]))
	case ERROR_ID_NOT_IN_REGION:
		response.Error(http.StatusBadRequest, fmt.Sprintf(`The %s region doesn't have a pokemon with ID %s`, strings.ToUpper(params["region"]), params["id"]))
	default:
		pkmn, err := ctrl.loadOne(id)
		if err != nil {
			if _, ok := err.(*PokemonNotFoundError); ok {
				response.Error(http.StatusNotFound, err.Error())
				break
			} else {
				response.Error(http.StatusInternalServerError, "The server encountered an error while processing your request. Please try again later")
				break
			}
		}
		response.Data = pkmn
		response.Status = http.StatusOK
	}
}

func (ctrl PokemonController) Update(params martini.Params) {

}

func (ctrl PokemonController) Delete(request *Request, params martini.Params) {
	id := regionalIDToNational(params["region"], params["id"])
	switch id {
	case ERROR_BAD_ID, ERROR_BAD_REGION, ERROR_ID_NOT_IN_REGION:
		request.Error(422, "Your request could not be completed as provided")
	default:
		var mons []Species
		ctrl.Database.Where("dex_number", "=", id).Limit(1).Find(&mons)
		_, err := ctrl.Database.Delete(&mons)
		if err != nil {
			request.Status = http.StatusNoContent
		} else {
			request.Error(http.StatusInternalServerError, "The server has encountered an error, please try again later")
		}
	}
}

func (ctrl PokemonController) Metadata(request *Request) {
	options := make(map[string]interface{})
	methods := make(map[string]interface{})

	methods["GET"] = map[string]interface{}{
		"url": "/:region/:id",
		"args": map[string]string{
			"region": "The region the pokemon is registered to",
			"id":     "The pokemon's numerical ID in the region specified",
		},
		"argument_types": map[string]string{
			"region": "national|kanto|johto|hoenn|sinnoh|unova|kalos",
			"id":     "Three-digit integer - ex. 001",
		},
	}

	options["resource"] = "/pokemon"
	options["methods"] = methods

	request.Data = options
	request.Status = http.StatusOK
}

func regionalIDToNational(region string, id string) int {
	nationalID, err := strconv.Atoi(id)

	if err != nil {
		return ERROR_BAD_ID
	}

	switch region {
	case "national":
		break
	case "kanto":
		if 0 < nationalID && nationalID < 152 {
			break
		} else {
			nationalID = ERROR_ID_NOT_IN_REGION
		}
	case "johto":
		nationalID += 151
	case "hoenn":
		nationalID += 251
	case "sinnoh":
		nationalID += 386
	case "unova":
		nationalID += 493
	case "kalos":
		nationalID += 649
	}
	return nationalID
}

func (ctrl PokemonController) loadOne(id int) (Species, error) {
	var results []Species
	err := ctrl.Database.Where("dex_number", "=", id).Limit(1).Find(&results)
	if err != nil {
		return Species{}, err
	}
	if len(results) < 1 {
		return Species{}, &PokemonNotFoundError{id}
	}
	return results[0], nil
}
