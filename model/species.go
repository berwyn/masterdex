package model

import (
	"github.com/eaigner/hood"
	"github.com/martini-contrib/binding"
	"net/http"
)

type Species struct {
	Id        hood.Id `form:"-" json:"-"`
	Name      string  `sql:"size(255),notnull" form:"name" json:"name" binding:"required"`
	DexNumber int     `sql:"notnull" form:"dex_number" json:"dex_number" binding:"required"`
}

func (pkmn Species) Validate(errors *binding.Errors, req *http.Request) {
	if pkmn.Name == "" {
		errors.Fields["name"] = "You must provide a name"
	}

	if pkmn.DexNumber == 0 {
		errors.Fields["dex_number"] = "You must provide an ID number > 0"
	}
}
