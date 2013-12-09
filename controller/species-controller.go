package controller

import (
	"github.com/eaigner/hood"
)

type SpeciesController struct {
	Database *hood.Hood
	create   interface{} `restr:"POST,Creates a new Pokemon species"`
}
