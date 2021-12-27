package models

import (
	"testing"

	"github.com/urlshortner/src/entity"
)

func TestEncodeUrlAndStore(t *testing.T) {
	type args struct {
		input entity.URL
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success : Encode URL", args{entity.URL{URL: "https://test.com/hello"}}, "http://urlshortner.com/oXhWr3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeUrlAndStore(tt.args.input); got != tt.want {
				t.Errorf("EncodeUrlAndStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
