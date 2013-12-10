package controller

import (
	. "../model"
	"errors"
	"fmt"
	"github.com/eaigner/hood"
	"log"
)

type SpeciesController struct {
	Database *hood.Hood
	create   interface{} `rest:"POST"`
	read     interface{} `rest:"GET"`
	update   interface{} `rest:"PUT"`
}

func (ctrl SpeciesController) Create(id string) (interface{}, error) {
	return nil, nil
}

func (ctrl SpeciesController) Read(id string) (interface{}, error) {
	var results []Species
	err := ctrl.Database.Where("dex_number", "=", id).Limit(1).Find(&results)
	if err != nil {
		// TODO Send a real error up
		return Species{}, errors.New("Database error")
	}
	return results[0], nil
}

func (ctrl SpeciesController) Update(id string) (interface{}, error) {
	return nil, nil
}
