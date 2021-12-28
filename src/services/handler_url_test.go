package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/urlshortner/src/entity"
	"github.com/urlshortner/src/tests"
)

var (
	router *mux.Router
)

func TestHandlePostEncodeUrl(t *testing.T) {
	url := entity.URL{URL: "https://twitter.com/tweet/getTweet/"}
	body, _ := json.Marshal(&url)
	router = startTestServer()
	t.Run("CreateShortenUrlSuccess", func(t *testing.T) {
		tests.ExecuteAndParseJSON(t, router, http.MethodPost, "/urlshortner/api/create", string(body), http.StatusCreated, url)
	})
	t.Run("CreateShortenUrlDecodeRequestFailed", func(t *testing.T) {
		tests.ExecuteAndParseJSON(t, router, http.MethodPost, "/urlshortner/api/create", "", http.StatusInternalServerError, url)
	})
	t.Run("CreateShortenUrlFailure", func(t *testing.T) {
		encodeUrlAndStore = func(input entity.URL) (string, error) {
			return "", errors.New("Failed to create/store shorten URL")
		}
		tests.ExecuteAndParseJSON(t, router, http.MethodPost, "/urlshortner/api/create", string(body), http.StatusInternalServerError, url)
	})
	os.Remove("data.txt")
}

func startTestServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/urlshortner/api/create", HandlePostEncodeUrl)
	return r
}
