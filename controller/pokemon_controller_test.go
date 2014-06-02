package controller_test

import (
	. "github.com/berwyn/masterdex"
	. "github.com/berwyn/masterdex/controller"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	controller PokemonController
	bulbasaur  Pokemon
	ivysaur    Pokemon
	data       map[int]Pokemon
)

type MockDataStore struct {
	Data map[int]Pokemon
}

func (store MockDataStore) Find(id int) interface{} {
	return store.Data[id]
}

func (store MockDataStore) FindAll() interface{} {
	values := make([]Pokemon, 0, len(store.Data))
	for _, value := range store.Data {
		values = append(values, value)
	}
	return values
}

var _ = Describe("PokemonController", func() {
	BeforeSuite(func() {
		bulbasaur = Pokemon{1, "Bulbasaur", "Seed Pokémon"}
		ivysaur = Pokemon{2, "Ivysaur", "Seed Pokemon"}
	})

	BeforeEach(func() {
		data = make(map[int]Pokemon)
		data[bulbasaur.DexNumber] = bulbasaur
		data[ivysaur.DexNumber] = ivysaur

		store := MockDataStore{data}
		controller = PokemonController{store}
	})

	Context("Successful requests", func() {
		It("should return a list of pokemon for the index", func() {
			req, _ := http.NewRequest("GET", "localhost:1234/pokemon", nil)
			res := httptest.NewRecorder()

			controller.Index(res, req)

			payload, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())
			Expect(payload).NotTo(BeNil())

			var result []Pokemon
			json.Unmarshal(payload, &result)
			Expect(result).NotTo(BeNil())
			Expect(len(result)).To(Equal(2))
			for _, value := range result {
				Expect(value).To(BeEquivalentTo(data[value.DexNumber]))
			}
		})

		It("should return a pokemon", func() {
			Fail("Not implemented")
		})

		It("should create a new pokemon", func() {
			Fail("Not Implemented")
		})

		It("should update a pokemon", func() {
			Fail("Not Implemented")
		})

		It("should delete a pokemon", func() {
			Fail("Not Implemented")
		})
	})

	Context("Failed requests", func() {

	})
})
