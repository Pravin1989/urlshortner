package utilities

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/urlshortner/src/entity"
)

func TestEncodeResponse(t *testing.T) {
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		r      *http.Request
		w      http.ResponseWriter
		resObj interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"Success : Encode", args{req, httptest.NewRecorder(), entity.URL{}}},
		{"Failure : Encode", args{req, httptest.NewRecorder(), make(chan (int))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EncodeResponse(tt.args.r, tt.args.w, tt.args.resObj)
		})
	}
}

func TestDecodeRequest(t *testing.T) {
	jsonStr, err := json.Marshal(entity.URL{URL: "https://test.com/test123"})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		r *http.Request
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Success : Decode", args{req, entity.URL{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodeRequest(tt.args.r, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("DecodeRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
