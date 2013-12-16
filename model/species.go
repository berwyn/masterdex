package model

import (
	"github.com/eaigner/hood"
)

type Species struct {
	Id        hood.Id `json:"-"`
	Name      string  `sql:"size(255),notnull" json:"name"`
	DexNumber int     `sql:"notnull" json:"dex_number"`
}
