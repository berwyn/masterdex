package controller_test

import (
	. "github.com/berwyn/masterdex"
	. "github.com/berwyn/masterdex/controller"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var (
	server     *Server
	controller PokemonController
	bulbasaur  Pokemon
	ivysaur    Pokemon
)

type MockDataStore struct {
	data map[int]Pokemon
}

func (store MockDataStore) Find(id int) interface{} {
	return store.data[id]
}

func (store MockDataStore) FindAll() interface{} {
	values := make([]Pokemon, len(store.data))
	for _, value := range store.data {
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
		data := make(map[int]Pokemon)
		data[bulbasaur.DexNumber] = bulbasaur
		data[ivysaur.DexNumber] = ivysaur

		store := MockDataStore{data}

		controller = PokemonController{store}
		server = NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("Successful requests", func() {
		It("should return a list of pokemon for the index", func() {
			server.AppendHandlers(
				controller.Index,
			)

			res, err := http.Get(server.URL())
			defer res.Body.Close()
			Expect(err).To(BeNil())
			Expect(res).NotTo(BeNil())

			payload, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())
			Expect(payload).NotTo(BeNil())

			fmt.Println(fmt.Sprintf("%v", payload))

			var result []Pokemon
			json.Unmarshal(payload, &result)
			Expect(len(result)).To(Equal(2))
		})

		It("should return a pokemon", func() {
			Fail("Not Implemented")
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
