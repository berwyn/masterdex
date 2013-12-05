package main

import (
	"github.com/berwyn/masterdex/model"
	"github.com/eaigner/hood"
)

func (m *M) CreateSpeciesTable_1386209731_Up(hd *hood.Hood) {
	hd.CreateTable(&model.Species{})
}

func (m *M) CreateSpeciesTable_1386209731_Down(hd *hood.Hood) {
	// TODO: implement
}
