package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/urlshortner/src/services"
)

const (
	port = 8090
)

var (
	handlePostEncodeUrl = handlers.HandlePostEncodeUrl
	handleGetEncodeUrl  = handlers.HandleOpenEncodedUrl
)

//This method add routes and start server
func Start() {
	router := loadRoutes()
	log.Println("Starting REST Server")
	log.Printf("REST server listening on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	fmt.Println("Server Crashed", err)
}

//Creates router
func loadRoutes() *mux.Router {
	serviceRouter := mux.NewRouter().PathPrefix("/urlshortner/api").Subrouter()
	registerApiRoutes(serviceRouter)
	return serviceRouter
}

//Regiter API methods inside router
func registerApiRoutes(r *mux.Router) {
	r.HandleFunc("/create", handlePostEncodeUrl).Methods(http.MethodPost)
	r.HandleFunc("/{uniqueId}", handleGetEncodeUrl).Methods(http.MethodGet)
}
