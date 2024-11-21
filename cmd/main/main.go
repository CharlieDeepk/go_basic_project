package main

import (
	"log"
	"net/http"

	"github.com/CharlieDeepk/go_basic_project/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))

}
