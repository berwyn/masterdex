package db

import (
	"github.com/eaigner/hood"
)

type Species struct {
	Id        hood.Id
	Name      string `sql:"size(255),notnull"`
	DexNumber int    `sql:"notnull"`
}

func (table *Species) Indexes(indexes *hood.Indexes) {
	indexes.AddUnique("dex_index", "dex_number")
}
