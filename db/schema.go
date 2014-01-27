package db

import (
	"github.com/eaigner/hood"
)

type Species struct {
	Id        hood.Id `json:"-"`
	Name      string  `sql:"size(255),notnull" json:"name"`
	DexNumber int     `sql:"notnull" json:"dex_number"`
}

func (table *Species) Indexes(indexes *hood.Indexes) {
	indexes.AddUnique("dex_index", "dex_number")
}
