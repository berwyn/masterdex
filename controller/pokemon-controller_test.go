package controller

import (
	. "github.com/berwyn/masterdex/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var (
	bulbasaur = Species{
		Name:      "Bulbasaur",
		DexNumber: 001,
	}
)

type MockPokmeonDatastore struct{}

func (MockPokmeonDatastore) Find(id string) (interface{}, error) {
	return bulbasaur, nil
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

		It("should reject IDs not in a region", func() {
			controller.Read(map[string]string{"region": "kanto", "id": "256"}, &request)

			Expect(request.Status).To(Equal(http.StatusBadRequest))
		})
	})
})
