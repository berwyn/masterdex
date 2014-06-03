package controller

import (
	. "github.com/berwyn/masterdex"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PokemonController struct {
	Data DataStore
}

func (controller *PokemonController) Register(server *mux.Router) {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).
		Methods("GET").
		Name("Pokemon Index")
	r.HandleFunc("/{id:[0-9]+}", controller.Get).
		Methods("GET").
		Name("Pokemon Get")
	server.Handle("/pokemon", r)
}

func (controller *PokemonController) Index(w http.ResponseWriter, r *http.Request) {
	entities := controller.Data.FindAll()
	payload, err := json.Marshal(&entities)
	if err == nil {
		fmt.Fprintf(w, string(payload[:]))
	} else {
		fmt.Fprintf(w, `{"error":500}`)
	}
}

func (controller *PokemonController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)

	if err != nil {
		fmt.Fprintf(w, `{"error":422}`)
		return
	}

	entity := controller.Data.Find(id)
	payload, err := json.Marshal(&entity)
	if err == nil {
		fmt.Fprintf(w, string(payload[:]))
	} else {
		fmt.Fprintf(w, `{"error":500}`)
	}
}

func (controller *PokemonController) Create(w http.ResponseWriter, r *http.Request) {

}

func (controller *PokemonController) Update(w http.ResponseWriter, r *http.Request) {

}

func (controller *PokemonController) Delete(w http.ResponseWriter, r *http.Request) {

}
