package model

import (
	"github.com/eaigner/hood"
)

type Species struct {
	Id        hood.Id
	Name      string `sql:"size(255),notnull"`
	DexNumber int    `sql:"notnull"`
}
