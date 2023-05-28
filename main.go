package main

import (
	"net/http"

	"github.com/XavierCZ-A/Rest-Api-Go/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", routes.IndexRoutes)
	http.ListenAndServe(":3000", r)

}
