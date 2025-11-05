package Util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
)

func SendJSON(status int, data any, w http.ResponseWriter) error {
	w.WriteHeader(status)
	deserialized, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(deserialized)
	return err
}

func DeserializeRequest(ptr any, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(ptr)
	return err
}

func Hash256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", string(h.Sum(nil)))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// CurrentDate returns the current date in YYYY-MM-DD format.
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}