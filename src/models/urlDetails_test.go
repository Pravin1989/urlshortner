package models

import (
	"os"
	"testing"

	"github.com/urlshortner/src/entity"
)

func TestEncodeUrlAndStore(t *testing.T) {
	type args struct {
		input entity.URL
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantErr   bool
		setUpMock func()
	}{
		{"Success : Encode URL", args{entity.URL{URL: "https://test.com/hello"}}, "http://urlshortner.com/oXhWr3", false, func() {
			readFile = func(fileName string) ([]byte, error) {
				return nil, nil
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setUpMock()
			got, err := EncodeUrlAndStore(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeUrlAndStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeUrlAndStore() = %v, want %v", got, tt.want)
			}
		})
		os.Remove("data.txt")
	}
}
