package models

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urlshortner/src/entity"
)

const (
	baseURL  = "http://urlshortner.com/"
	filename = "data.txt"
)

var (
	readFile = func(filename string) ([]byte, error) {
		return ioutil.ReadFile(filename)
	}
)

func EncodeUrlAndStore(input entity.URL) (string, error) {
	// value, isExist := urlsMap[input.URL]
	urlsMap, readErr := ReadFromFile()
	if readErr != nil {
		return "", readErr
	}
	value, isExist := urlsMap[input.URL]
	if isExist {
		return value, nil
	}
	encodedValueToStore := Encode(input)
	urlsMap[encodedValueToStore] = input.URL
	if writeErr := WriteToFile(urlsMap); writeErr != nil {
		return "", writeErr
	}
	return baseURL + encodedValueToStore, nil
}

func Encode(input entity.URL) string {
	hasher := sha1.New()
	hasher.Write([]byte(input.URL))
	encodedValue := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encodedValue[0:6]
}

func WriteToFile(data map[string]string) error {
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
	n, writeError := file.WriteString(string(userData))
	if writeError != nil {
		log.Printf("failed writing to file: %v", writeError)
		return writeError
	}
	fmt.Println("Length :", n)
	defer file.Close()
	return nil
}

func ReadFromFile() (map[string]string, error) {
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
