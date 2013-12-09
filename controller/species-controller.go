package controller

import (
	"github.com/eaigner/hood"
)

type SpeciesController struct {
	Database *hood.Hood
	create   interface{} `rest:"POST"`
	update   interface{} `rest:"PUT"`
}

func (ctrl SpeciesController) Create(id string) (interface{}, error) {
	return nil, nil
}

func (ctrl SpeciesController) Update(id string) (interface{}, error) {
	return nil, nil
}
