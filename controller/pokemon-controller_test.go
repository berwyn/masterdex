package controller

import (
	"errors"
	. "github.com/berwyn/masterdex/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"strconv"
)

var (
	bulbasaur = Pokemon{Name: "Bulbasaur", DexNumber: 1}
	ivysaur   = Pokemon{Name: "Ivysaur", DexNumber: 2}
	pkmnData  = map[string]Pokemon{
		"1": bulbasaur,
		"2": ivysaur,
	}
)

type MockPokmeonDatastore struct{
	MockDatastore
}

func (MockPokmeonDatastore) Find(id string) (interface{}, error) {
	if pkmn, ok := pkmnData[id]; ok {
		return pkmn, nil
	}
	return Pokemon{}, &PokemonNotFoundError{id}
}

func (MockPokmeonDatastore) Insert(entity interface{}) (interface{}, error) {
	return entity, nil
}

func (MockPokmeonDatastore) Update(entity interface{}) (interface{}, error) {
	if pkmn, ok := entity.(Pokemon); ok {
		var idString = strconv.Itoa(pkmn.DexNumber)
		if _, ok := pkmnData[idString]; ok {
			return entity, nil
		}
		return Pokemon{}, &PokemonNotFoundError{idString}
	}
	return Pokemon{}, errors.New("You didn't provide a pokemon")
}

func (MockPokmeonDatastore) Delete(id string) error {
	if _, ok := pkmnData[id]; ok {
		return nil
	}
	return &PokemonNotFoundError{id}
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

	Describe("Helper functions", func() {
		It("should accept all national IDs", func() {
			result := regionalIDToNational("national", "222")

			Expect(result).To(Equal(222))
		})

		It("should reject non-integer IDs", func() {
			floatID := regionalIDToNational("national", "72.3")
			stringID := regionalIDToNational("national", "xzy")

			Expect(floatID).To(Equal(ERROR_BAD_ID))
			Expect(stringID).To(Equal(ERROR_BAD_ID))
		})

		It("should reject invalid regions", func() {
			badRegion := regionalIDToNational("xzy", "222")

			Expect(badRegion).To(Equal(ERROR_BAD_REGION))
		})

		It("should reject Kanto IDs above 151", func() {
			result := regionalIDToNational("kanto", "152")

			Expect(result).To(Equal(ERROR_ID_NOT_IN_REGION))
		})
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
		It("should accept a payload", func() {
			controller.Create(ivysaur, &request)

			Expect(request.Status).To(Equal(http.StatusCreated))
			Expect(request.Data).To(BeEquivalentTo(&ivysaur))
		})
	})

	Describe("PUT", func() {
		It("should accept a payload", func() {
			ivysaurCopy := ivysaur
			ivysaurCopy.Name = "IvysaurCopy"

			controller.Update(ivysaurCopy, &request)

			Expect(request.Status).To(Equal(http.StatusOK))
			Expect(request.Data).To(BeEquivalentTo(ivysaurCopy))
		})

		It("should create pokemon that don't exist", func() {
			venusaur := Pokemon{
				Name:      "Venusaur",
				DexNumber: 3,
			}

			controller.Create(venusaur, &request)

			Expect(request.Status).To(Equal(http.StatusCreated))
			Expect(request.Data).To(BeEquivalentTo(&venusaur))
		})
	})

	Describe("PATCH", func() {
		It("should reject pokemon that don't exist", func() {
			venusaur := Pokemon{Name: "Venusaur", DexNumber: 3}

			controller.Update(venusaur, &request)

			Expect(request.Status).To(Equal(http.StatusNotFound))
		})
	})

	Describe("DELETE", func() {
		It("should accept delete requests for existing pokemon", func() {
			controller.Delete(map[string]string{"dex": "national", "id": "2"}, &request)

			Expect(request.Status).To(Equal(http.StatusNoContent))
		})

		It("should reject delete requests for pokemon that don't exist", func() {
			controller.Delete(map[string]string{"dex": "national", "id": "5"}, &request)

			Expect(request.Status).To(Equal(http.StatusNotFound))
		})
	})
})
