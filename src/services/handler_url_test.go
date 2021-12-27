package handlers

import (
	"encoding/json"
	"net/http"
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
	t.Run("CreteShortenUrlSuccess", func(t *testing.T) {
		tests.ExecuteAndParseJSON(t, router, http.MethodPost, "/urlshortner/api/create", string(body), http.StatusCreated, url)
	})
	t.Run("CreteShortenUrlFailure", func(t *testing.T) {
		tests.ExecuteAndParseJSON(t, router, http.MethodPost, "/urlshortner/api/create", "", http.StatusInternalServerError, url)
	})
}

func startTestServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/urlshortner/api/create", HandlePostEncodeUrl)
	return r
}
