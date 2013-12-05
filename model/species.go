package model

import (
	"errors"
	"github.com/eaigner/hood"
)

type Species struct {
	Id        hood.Id
	Name      string `sql:"size(255),notnull"`
	DexNumber int    `sql:"notnull"`
}

func LoadSpecies(db *hood.Hood, dex string, id string) (Species, error) {
	if dex == "national" {
		var queryResults []Species
		db.Where("dex_number", "=", id).Limit(1).Find(&queryResults)
		if len(queryResults) >= 1 {
			return queryResults[0], nil
		} else {
			return Species{}, errors.New("Not Found")
		}
	}
	return Species{}, errors.New("NYI")
}
