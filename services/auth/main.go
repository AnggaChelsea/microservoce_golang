package main

import(
	"simplelogic/services/auth/repository/entity"
	"fmt"
	"simplelogic/services/auth/repository/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.AuthMiddleware(http.HandlerFunc(entity.Demo1API))).Methods("GET")
	err := http.ListenAndServe(":8080", router); if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server is running on port 8080")
}