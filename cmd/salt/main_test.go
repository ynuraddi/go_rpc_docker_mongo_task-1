package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// generateRandomSalt
func TestGenerateRandomSalt(t *testing.T) {
	cases := []struct {
		name string
		want func(string) bool
	}{
		{
			name: "check len of salt",
			want: func(s string) bool {
				return len(s) == saltSize
			},
		},
		{
			name: "check charset of salt",
			want: func(s string) bool {
				for _, v := range s {
					if !strings.ContainsRune(charset, v) {
						return false
					}
				}

				return true
			},
		},
	}

	got := string(generateRandomSalt())

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			if !test.want(got) {
				t.Errorf("got %v", got)
			}
		})
	}
}

// TestHandler
func TestHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/create-user", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong Content-Type header: got %v want %v",
			contentType, expectedContentType)
	}

	var s s
	err = json.Unmarshal(rr.Body.Bytes(), &s)
	if err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	if len(s.Salt) != 12 {
		t.Errorf("handler returned invalid salt: got %v want %v",
			s.Salt, 12)
	}
}
