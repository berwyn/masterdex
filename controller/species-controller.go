package controller

import (
	"github.com/eaigner/hood"
)

type SpeciesController struct {
	Database *hood.Hood
	create   interface{} `restr:"POST,Creates a new Pokemon species"`
}

func (ctrl *SpeciesController) Create(id string) (interface{}, error) {
	return nil, nil
}
