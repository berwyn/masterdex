package main

import (
	"encoding/json"
	"github.com/eaigner/hood"
	"io/ioutil"
	"log"
	_ "model"
	"net/http"
	"os"
)

const (
	port = ":1234"
)

var db *hood.Hood

func main() {

	// First, load our DB config
	config := loadDbConfig()
	db = openDatabase(config["development"]["driver"], config["development"]["source"])

	// Load our handlers
	http.HandleFunc("/dex/pokemon/", SpeciesHandler)

	// Start the server
	debugLog("Booting server on " + port)
	http.ListenAndServe(port, nil)
}

func debugLog(message string) {
	log.Println("[DEBUG]", message)
}

func warnLog(message string) {
	log.Println("[WARN]", message)
}

func errLog(message string) {
	log.Println("[ERROR]", message)
}

func loadDbConfig() (config map[string]map[string]string) {
	file, err := ioutil.ReadFile("db/config.json")
	if err != nil {
		log.Fatal("Couldn't open database config")
		os.Exit(1)
	}

	if err = json.Unmarshal(file, &config); err != nil {
		log.Fatal("Couldn't deserialise database config")
		os.Exit(1)
	}
	return config
}

func openDatabase(driver string, connectionString string) (database *hood.Hood) {
	database, err := hood.Open(driver, connectionString)
	if err != nil {
		log.Fatal("Couldn't connect to database")
		os.Exit(1)
	}
	return database
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO Serve 404s and whatnot
		fn(res, req)
	}
}

func SpeciesHandler(res http.ResponseWriter, req *http.Request) {

}
