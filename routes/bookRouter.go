package routes

import (
	"simplelogic/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/buku", controllers.AmbilSemuaBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controllers.AmbilBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku", controllers.TmbhBuku).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controllers.UpdateBuku).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controllers.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}