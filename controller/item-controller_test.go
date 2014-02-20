package controller

import (
	. "github.com/onsi/ginkgo"
)

type MockItemDataStore struct{}

func (MockItemDataStore) Find(id string) (interface{}, error) {
	return new(interface{}), nil
}

func (MockItemDataStore) Insert(entity interface{}) (interface{}, error) {
	return new(interface{}), nil
}

func (MockItemDataStore) Update(entity interface{}) (interface{}, error) {
	return new(interface{}), nil
}

func (MockItemDataStore) Delete(id string) error {
	return nil
}

var _ = Describe("Item Controller", func() {

	var (
		controller ItemController
	)

	BeforeEach(func() {
		controller = ItemController{MockItemDataStore{}}
	})

})
