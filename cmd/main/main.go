package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/melvin-n/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	//passing routes defined in routes folder
	routes.RegisterBookStoreRoutes(r)
	//http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
