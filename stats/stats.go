package stats

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stats", Stats).Methods("GET")
	return router
}

func Stats(w http.ResponseWriter, r *http.Request) {

}
