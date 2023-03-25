package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type s struct {
	Salt string `json:"salt"`
}

func main() {
	http.HandleFunc("/generate-salt", handler)

	http.ListenAndServe(":8080", nil)
}

const (
	saltSize = 12
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateRandomSalt() []byte {
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, saltSize)

	for i := 0; i < saltSize; i++ {
		salt[i] = byte(charset[rand.Intn(len(charset))])
	}

	return salt
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	s := s{}

	s.Salt = string(generateRandomSalt())

	data, err := json.Marshal(s)
	if err != nil {
		// logger
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
