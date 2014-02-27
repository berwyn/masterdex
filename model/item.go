package model

import (
	"github.com/eaigner/hood"
)

type Item struct {
	Id       hood.Id `form:"-" json:"-"`
	Name     string  `form:"name" json:"name"`
	SlugName string  `form:"slug" json:"slug"`
}
