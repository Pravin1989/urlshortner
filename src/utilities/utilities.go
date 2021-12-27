package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeResponse(r *http.Request, w http.ResponseWriter, resObj interface{}) {
	err := json.NewEncoder(w).Encode(&resObj)
	if err != nil {
		fmt.Println("Error Occured while encoding response: ", err)
	}
}

func DecodeRequest(r *http.Request, i interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&i)
}
