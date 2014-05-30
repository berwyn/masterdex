package masterdex

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	http.ListenAndServe(":1234", nil)
}
