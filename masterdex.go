package main

import (
	"encoding/json"
	"fmt"
	. "github.com/berwyn/masterdex/controller"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/eaigner/hood"
	"html/template"
	"io/ioutil"
	"net/http"
)

const (
	port = ":1234"
)

var db *hood.Hood
var controllers []Controller

func main() {
	// First, load our DB config
	config := loadDbConfig()
	db = openDatabase(config["development"]["driver"], config["development"]["source"])

	// Populate the controller slice
	controllers = []Controller{
		PokemonController{Database: db},
		ItemController{Database: db},
		AboutController{},
		RootController{},
	}

	// Load our handlers
	m := configureMartini()
	http.Handle("/", m)

	// Start the server
	debugLog("Booting server on " + port)
	http.ListenAndServe(port, nil)
}

func debugLog(message string) {
	fmt.Println("[masterdex][DEBUG]", message)
}

func warnLog(message string) {
	fmt.Println("[masterdex][WARN]", message)
}

func errLog(message string) {
	fmt.Println("[masterdex][ERROR]", message)
}

func loadDbConfig() (config map[string]map[string]string) {
	file, err := ioutil.ReadFile("db/config.json")
	if err != nil {
		panic("Couldn't open database config")
	}

	if err = json.Unmarshal(file, &config); err != nil {
		panic("Couldn't deserialise database config")
	}
	return config
}

func openDatabase(driver string, connectionString string) (database *hood.Hood) {
	database, err := hood.Open(driver, connectionString)
	if err != nil {
		panic("Couldn't connect to database")
	}
	return database
}

func configureMartini() *martini.ClassicMartini {
	m := martini.Classic()
	helpers := []template.FuncMap{}
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Directory:  "views",
		Extensions: []string{".html"},
		Funcs:      helpers,
	}))
	for _, ctrl := range controllers {
		ctrl.Register(m)
	}
	return m
}
