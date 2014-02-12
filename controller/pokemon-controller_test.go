package controller

import (
	"errors"
	. "github.com/berwyn/masterdex/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var (
	bulbasaur = Species{Name: "Bulbasaur", DexNumber: 1}
	venusaur  = Species{Name: "Venusaur", DexNumber: 2}
)

type MockPokmeonDatastore struct{}

func (MockPokmeonDatastore) Find(id string) (interface{}, error) {
	if id != "1" {
		return bulbasaur, &PokemonNotFoundError{id}
	}
	return bulbasaur, nil
}

func (MockPokmeonDatastore) Insert(entity interface{}) (interface{}, error) {
	if _, ok := entity.(*Species); ok {
		return entity, nil
	}
	return Species{}, errors.New("That isn't a species")
}

var _ = Describe("Pokemon controller", func() {

	var (
		request    Request
		controller PokemonController
	)

	BeforeEach(func() {
		request = Request{}
		controller = PokemonController{MockPokmeonDatastore{}}
	})

	It("should handle root requests", func() {
		controller.Index(&request)

		Expect(request.Status).To(Equal(http.StatusOK))
		Expect(request.Template).To(Equal("pokemon"))
	})

	Describe("GET", func() {
		It("should return a pokémon", func() {
			controller.Read(map[string]string{"region": "national", "id": "001"}, &request)

			Expect(request.Status).To(Equal(http.StatusOK))
			Expect(request.Data).To(Equal(bulbasaur))
			Expect(request.Template).To(Equal("pokemon"))
		})

		It("should reject bad IDs", func() {
			controller.Read(map[string]string{"region": "national", "id": "q"}, &request)

			Expect(request.Status).To(Equal(http.StatusBadRequest))
		})

		It("should reject bad regions", func() {
			controller.Read(map[string]string{"region": "xzy", "id": "001"}, &request)

			Expect(request.Status).To(Equal(http.StatusBadRequest))
		})

		It("should reject IDs not in a region", func() {
			controller.Read(map[string]string{"region": "kanto", "id": "256"}, &request)

			Expect(request.Status).To(Equal(http.StatusBadRequest))
		})

		It("should return an error for pokemon that don't exist", func() {
			controller.Read(map[string]string{"region": "national", "id": "151"}, &request)

			Expect(request.Status).To(Equal(http.StatusNotFound))
		})
	})

	Describe("POST", func() {
		Context("JSON payloads", func() {
			BeforeEach(func() {
				request.UsingJSON = true
				request.ContainsJSON = true
			})

			It("should accept a payload", func() {
				request.Payload = []byte(`{"name":"Venusaur","dex_number":2}`)

				controller.Create(&request)

				Expect(request.Status).To(Equal(http.StatusCreated))
				Expect(request.Data).To(BeEquivalentTo(&venusaur))
			})

			It("should reject malformed payloads", func() {
				request.Payload = []byte(`{"name":Venusaur,"dex_number"2`)

				controller.Create(&request)

				Expect(request.Status).To(Equal(422))
			})

			It("should reject payloads that aren't pokemon", func() {
				request.Payload = []byte(`{"hobbies":"eating cheese"}`)

				controller.Create(&request)

				Expect(request.Status).To(Equal(422))
			})
		})
	})
})
