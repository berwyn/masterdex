package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateSpeciesDexIndex_1386209789_Up(hd *hood.Hood) {
	hd.CreateIndex("species", "dex_index", true, "dex_number")
}

func (m *M) CreateSpeciesDexIndex_1386209789_Down(hd *hood.Hood) {
	// TODO: implement
}
