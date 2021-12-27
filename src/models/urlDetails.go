package models

import (
	"github.com/urlshortner/src/entity"
)

var urlsMap = make(map[int]interface{})

func EncodeUrlAndStore(input entity.URL) string {
	value, isExist := urlsMap[len(input.URL)]
	if isExist {
		return value.(string)
	}
	urlsMap[len(input.URL)] = input.URL
	return input.URL
}
