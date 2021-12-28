package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Encode Object to JSON
func EncodeResponse(r *http.Request, w http.ResponseWriter, resObj interface{}) {
	err := json.NewEncoder(w).Encode(&resObj)
	if err != nil {
		fmt.Println("Error Occured while encoding response: ", err)
	}
}

// Decode JSON request to Object
func DecodeRequest(r *http.Request, i interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&i)
}
