package main

import (
	"log"
	"net/http"

	"github.com/arniemutasa/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	//start server
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
