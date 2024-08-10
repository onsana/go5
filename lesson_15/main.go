package main

import (
	"net/http"

	"library-server/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	http.ListenAndServe(":8000", router)
}
