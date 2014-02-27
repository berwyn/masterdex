package controller

import (
	. "github.com/berwyn/masterdex/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var (
	potion      = Item{Name: "Potion", SlugName: "potion"}
	superPotion = Item{Name: "Super Potion", SlugName: "super-potion"}
	itemData    = map[string]Item{
		potion.SlugName: potion,
	}
)

type MockItemDataStore struct{}

func (MockItemDataStore) Find(id string) (interface{}, error) {
	if item, ok := itemData[id]; ok {
		return item, nil
	}
	return Item{}, &ItemNotFoundError{id}
}

func (MockItemDataStore) Insert(entity interface{}) (interface{}, error) {
	if _, ok := itemData[entity.(Item).SlugName]; !ok {
		return entity, nil
	}
	return Item{}, &ItemValidationError{}
}

func (MockItemDataStore) Update(entity interface{}) (interface{}, error) {
	return new(interface{}), nil
}

func (MockItemDataStore) Delete(id string) error {
	return nil
}

var _ = Describe("Item Controller", func() {

	var (
		request    Request
		controller ItemController
	)

	BeforeEach(func() {
		request = Request{}
		controller = ItemController{MockItemDataStore{}}
	})

	It("should handle root requests", func() {
		controller.Index(&request)

		Expect(request.Status).To(Equal(http.StatusOK))
		Expect(request.Template).To(Equal("item"))
	})

	Describe("GET", func() {
		It("should return an item", func() {
			controller.Read(map[string]string{"slug": potion.SlugName}, &request)

			Expect(request.Status).To(Equal(http.StatusOK))
			Expect(request.Template).To(Equal("item"))
			Expect(request.Data).To(BeEquivalentTo(potion))
		})

		It("should reject slugs that don't exist", func() {
			controller.Read(map[string]string{"slug": "abracadabra"}, &request)

			Expect(request.Status).To(Equal(http.StatusNotFound))
		})
	})

	Describe("POST", func() {
		It("should accept payloads", func() {
			controller.Create(superPotion, &request)

			Expect(request.Status).To(Equal(http.StatusCreated))
			Expect(request.Data).To(BeEquivalentTo(superPotion))
			Expect(request.Template).To(Equal("item"))
		})

		It("should reject existing items", func() {
			controller.Create(potion, &request)

			Expect(request.Status).To(Equal(422))
		})
	})
})
