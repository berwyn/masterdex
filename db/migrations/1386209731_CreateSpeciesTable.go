package main

import (
	"github.com/eaigner/hood"
)

type Pokemon struct {
	Id        hood.Id `json:"-"`
	Name      string  `sql:"size(255),notnull" json:"name"`
	DexNumber int     `sql:"notnull" json:"dex_number"`
}

func (m *M) CreatePokemonTable_1386209731_Up(hd *hood.Hood) {
	hd.CreateTable(&Pokemon{})
}

func (m *M) CreatePokemonTable_1386209731_Down(hd *hood.Hood) {
	// TODO: implement
}
