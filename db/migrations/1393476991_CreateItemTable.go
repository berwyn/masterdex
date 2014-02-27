package main

import (
	"github.com/eaigner/hood"
)

type Item struct {
	Id       hood.Id `form:"-" json:"-"`
	Name     string  `form:"name" json:"name"`
	SlugName string  `form:"slug" json:"slug"`
}

func (m *M) CreateItemTable_1393476991_Up(hd *hood.Hood) {
	hd.CreateTable(&Item{})
}

func (m *M) CreateItemTable_1393476991_Down(hd *hood.Hood) {
	hd.DropTable(&Item{})
}
