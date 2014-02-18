package controller

import (
	"fmt"
	. "github.com/berwyn/masterdex/model"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"net/http"
	"strconv"
	"strings"
)

// These define our basic error states
const (
	// The ID is completely invalid
	ERROR_BAD_ID = -1
	// The region is completely invalid
	ERROR_BAD_REGION = -2
	// The ID doesn't exist for the specified region
	ERROR_ID_NOT_IN_REGION = -3
)

// Error raised when a pokémon can't be found
type PokemonNotFoundError struct {
	id string
}

func (err PokemonNotFoundError) Error() string {
	return fmt.Sprintf("Could not find the pokemon with ID %d", err.id)
}

// Martini controller for /pokemon
type PokemonController struct {
	datastore Datastore
}

// Register all our valid endpoints
func (ctrl PokemonController) Register(server *martini.ClassicMartini) {
	server.Get("/pokemon", ctrl.Index)
	server.Get("/pokemon/:dex/:id", ctrl.Read)
	server.Post("/pokemon", binding.Bind(Pokemon{}), binding.ErrorHandler, ctrl.Create)
	server.Put("/pokemon/:dex/:id", binding.Bind(Pokemon{}), binding.ErrorHandler, ctrl.Update)
	server.Patch("/pokemon/:dex/:id", binding.Bind(Pokemon{}), binding.ErrorHandler, ctrl.Update)
	server.Delete("/pokemon/:dex/:id", ctrl.Delete)
	server.Options("/pokemon", ctrl.Metadata)
}

// Load HTML, or for JSON requests return the same as OPTIONS
func (ctrl PokemonController) Index(response *Request) {
	if response.UsingJSON {
		ctrl.Metadata(response)
	} else {
		response.Status = 200
		response.Data = new(struct{})
		response.Template = "pokemon"
	}
}

func (ctrl PokemonController) Create(pkmn Pokemon, response *Request) {
	entity, err := ctrl.datastore.Insert(&pkmn)
	if err != nil {
		response.Error(http.StatusInternalServerError, "We couldn't save that Pokemon, please try again later")
	}
	response.Data = entity
	response.Status = http.StatusCreated
	response.Template = "pokemon"
}

func (ctrl PokemonController) Read(params martini.Params, response *Request) {
	id := regionalIDToNational(params["region"], params["id"])
	switch id {
	case ERROR_BAD_ID:
		response.Error(
			http.StatusBadRequest,
			fmt.Sprintf(`'%s' is not a valid ID for a pokemon`, params["id"]),
		)
	case ERROR_BAD_REGION:
		response.Error(
			http.StatusBadRequest,
			fmt.Sprintf(`'%s' is not a valid region`, params["region"]),
		)
	case ERROR_ID_NOT_IN_REGION:
		response.Error(
			http.StatusBadRequest,
			fmt.Sprintf(`The %s region doesn't have a pokemon with ID %s`, strings.ToUpper(params["region"]), params["id"]),
		)
	default:
		pkmn, err := ctrl.datastore.Find(strconv.Itoa(id))
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
		response.Template = "pokemon"
		response.Status = http.StatusOK
	}
}

func (ctrl PokemonController) Update(payload Pokemon, reqeust *Request) {
	entity, err := ctrl.datastore.Update(payload)
	if err != nil {
		if _, ok := err.(*PokemonNotFoundError); ok {
			reqeust.Error(http.StatusNotFound, "The pokemon you're trying to update doesn't exist")
			return
		}

		reqeust.Error(http.StatusInternalServerError, "There was an error handling your request, please try again later")
		return
	}
	reqeust.Data = entity
	reqeust.Status = http.StatusOK
	reqeust.Template = "pokemon"
}

func (ctrl PokemonController) Delete(params martini.Params, request *Request) {
	id := strconv.Itoa(regionalIDToNational(params["dex"], params["id"]))
	err := ctrl.datastore.Delete(id)
	if err != nil {
		if _, ok := err.(*PokemonNotFoundError); ok {
			request.Error(http.StatusNotFound, err.Error())
			return
		}
		request.Error(http.StatusInternalServerError, "There was an issue processing your request, please try again later")
		return
	}
	request.Status = http.StatusNoContent
	request.Template = "pokemon"
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

// This helper function turns a regional string ID into it's regions
// integer equivalent
func regionalIDToNational(region string, id string) int {
	nationalID, err := strconv.Atoi(id)

	if err != nil || nationalID < 0 {
		return ERROR_BAD_ID
	}

	switch region {
	case "national":
		break
	case "kanto":
		if nationalID > 151 {
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
	default:
		nationalID = ERROR_BAD_REGION
	}
	return nationalID
}
