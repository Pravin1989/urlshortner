package models

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/urlshortner/src/entity"
)

const (
	baseURL  = "http://localhost:8090/urlshortner/api/"
	filename = "data.txt"
)

var (
	readFile = func(filename string) ([]byte, error) {
		return ioutil.ReadFile(filename)
	}
)

// This method encode URL and store inside file
func EncodeUrlAndStore(input entity.URL) (string, error) {
	urlsMap, readErr := readFromFile()
	if readErr != nil {
		return "", readErr
	}
	value, isExist := urlsMap[input.URL]
	if isExist {
		return value, nil
	}
	encodedValueToStore, _ := uuid.NewUUID()
	urlsMap[encodedValueToStore.String()] = input.URL
	if writeErr := writeToFile(urlsMap); writeErr != nil {
		return "", writeErr
	}
	return baseURL + encodedValueToStore.String(), nil
}
func GetDecodedUrl(key string) (string, error, int) {
	urlsMap, readErr := readFromFile()
	if readErr != nil {
		return "", readErr, http.StatusInternalServerError
	}
	value, isExist := urlsMap[key]
	if isExist {
		return value, nil, http.StatusOK
	}
	return "", errors.New("The required url not found"), http.StatusNotFound
}

// This method encode URL
func encode(input entity.URL) string {
	hasher := sha1.New()
	hasher.Write([]byte(input.URL))
	encodedValue := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encodedValue[0:6]
}

// This method write to file
func writeToFile(data map[string]string) error {
	file, fileCreateError := os.Create(filename)
	if fileCreateError != nil {
		log.Printf("Failed creating file: %s", fileCreateError)
		return fileCreateError

	}
	userData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, writeError := file.WriteString(string(userData))
	if writeError != nil {
		log.Printf("failed writing to file: %v", writeError)
		return writeError
	}
	defer file.Close()
	return nil
}

// This method read from file
func readFromFile() (map[string]string, error) {
	data, readError := readFile(filename)
	myMap := make(map[string]string)
	if readError != nil {
		return myMap, nil
	}
	if len(data) == 0 {
		return myMap, nil
	}
	marshalError := json.Unmarshal(data, &myMap)
	if marshalError != nil {
		log.Printf("Failed to marshal: %s", marshalError)
		return nil, readError
	}
	return myMap, nil

}
