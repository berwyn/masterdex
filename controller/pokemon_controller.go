package controller

import (
	. "github.com/berwyn/masterdex"

	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonController struct {
	Data DataStore
}

func (controller *PokemonController) Register(router *http.ServeMux) {

}

func (controller *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	entities := controller.Data.FindAll()
	payload, err := json.Marshal(&entities)
	if err == nil {
		fmt.Fprintf(w, "%v", string(payload[:]))
	} else {
		fmt.Fprintf(w, `{"error":500}`)
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
