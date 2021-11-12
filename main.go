package main

import (
	"fmt"
	"log"
	"net/http"
	"simplelogic/routes"
)

func main() {
	r := routes.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)

	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
