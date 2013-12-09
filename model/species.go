package model

import (
	"github.com/eaigner/hood"
)

type Species struct {
	Id        hood.Id `json:"-"`
	Name      string  `sql:"size(255),notnull" json:"name"`
	DexNumber int     `sql:"notnull" json:"dex_number"`
}

func SpeciesDescriptor() map[string]string {
	return map[string]string{
		"GET":    "Returns the Pokemon identified by :id",
		"POST":   "Creates a new Pokemon based on the filled form [NYI]",
		"PUT":    "Updates a Pokemon identified by :id [NYI]",
		"PATCH":  "See PUT",
		"DELETE": "Deletes the Pokemon identified by :id",
	}
}
