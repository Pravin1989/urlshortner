package models

import (
	"crypto/sha1"
	"encoding/base64"

	"github.com/urlshortner/src/entity"
)

const (
	baseURL = "http://urlshortner.com/"
)

var urlsMap = make(map[string]string)

func EncodeUrlAndStore(input entity.URL) string {
	value, isExist := urlsMap[input.URL]
	if isExist {
		return value
	}
	encodedValueToStore := Encode(input)
	urlsMap[encodedValueToStore] = input.URL
	return baseURL + encodedValueToStore
}

func Encode(input entity.URL) string {
	hasher := sha1.New()
	hasher.Write([]byte(input.URL))
	encodedValue := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return encodedValue[0:6]
}
