package main

import (
	"log"
	"net/http"
	"github.com/davoodmossgreen/books-i/internal/routes"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}