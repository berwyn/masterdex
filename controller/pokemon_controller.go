package controller

import (
	. "github.com/berwyn/masterdex"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type PokemonController struct {
	Data DataStore
}

func (controller *PokemonController) Register(router *mux.Router) {

}

func (controller *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(controller.Data.FindAll())
	if err != nil {
		fmt.Fprintf(w, "%v", payload)
	}
}

func (controller *PokemonController) Get(w http.ResponseWriter, r *http.Request) {

}

func (controller *PokemonController) Create(w http.ResponseWriter, r *http.Request) {

}

func (controller *PokemonController) Update(w http.ResponseWriter, r *http.Request) {

}

func (controller *PokemonController) Delete(w http.ResponseWriter, r *http.Request) {

}
