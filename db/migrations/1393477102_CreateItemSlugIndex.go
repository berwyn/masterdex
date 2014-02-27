package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateItemSlugIndex_1393477102_Up(hd *hood.Hood) {
	hd.CreateIndex("items", "slug_index", true, "slug_name")
}

func (m *M) CreateItemSlugIndex_1393477102_Down(hd *hood.Hood) {
	hd.DropIndex("items", "slug_index")
}
