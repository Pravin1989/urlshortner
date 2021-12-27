package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/urlshortner/src/services"
)

var (
	handlePostEncodeUrl = handlers.HandlePostEncodeUrl
)

func Start() {
	router := loadRoutes()
	log.Println("Starting REST Server")
	log.Printf("REST server listening on port %d", 8090)
	err := http.ListenAndServe(":8090", router)
	fmt.Println("Server Crashed", err)
}

func loadRoutes() *mux.Router {
	serviceRouter := mux.NewRouter().PathPrefix("/urlshortner/api").Subrouter()
	registerApiRoutes(serviceRouter)
	return serviceRouter
}

func registerApiRoutes(r *mux.Router) {
	r.HandleFunc("/create", handlePostEncodeUrl).Methods(http.MethodPost)
}
