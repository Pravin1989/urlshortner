package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urlshortner/src/entity"
	"github.com/urlshortner/src/models"
	"github.com/urlshortner/src/utilities"
)

var (
	decodeRequest     = utilities.DecodeRequest
	encodeUrlAndStore = models.EncodeUrlAndStore
	decodeUrl         = models.GetDecodedUrl
)

// This is POST API method to create short url of long
func HandlePostEncodeUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started processing of HandlePostEncodeUrl")
	var url entity.URL
	decodeRequestError := decodeRequest(r, &url)
	if decodeRequestError != nil {
		log.Println("Invalid request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, writeErr := encodeUrlAndStore(url)
	if writeErr != nil {
		log.Println("Failed to write to file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utilities.EncodeResponse(r, w, &response)
	fmt.Println("Completed processing of HandlePostEncodeUrl")
}

func HandleOpenEncodedUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started processing of HandleOpenEncodedUrl")
	vars := mux.Vars(r)
	encodedUrl, isExist := vars["uniqueId"]
	if !isExist {
		log.Println("The parameter is missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value, readError, responseError := decodeUrl(encodedUrl)
	if readError != nil {
		log.Println("The error while reading :", readError)
		w.WriteHeader(responseError)
		return
	}
	http.Redirect(w, r, value, 200)
	// utilities.EncodeResponse(r, w, &value)
	fmt.Println("Completed processing of HandleOpenEncodedUrl")
}
