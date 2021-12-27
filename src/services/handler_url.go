package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urlshortner/src/entity"
	"github.com/urlshortner/src/models"
	"github.com/urlshortner/src/utilities"
)

var (
	decodeRequest     = utilities.DecodeRequest
	encodeUrlAndStore = models.EncodeUrlAndStore
)

func HandlePostEncodeUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started processing of HandlePostEncodeUrl")
	var url entity.URL
	decodeRequestError := decodeRequest(r, &url)
	if decodeRequestError != nil {
		log.Println("Invalid request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := encodeUrlAndStore(url)
	w.WriteHeader(http.StatusCreated)
	utilities.EncodeResponse(r, w, &response)
	fmt.Println("Completed processing of HandlePostEncodeUrl")
}
