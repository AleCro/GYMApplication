package main

import (
	"fmt"
	"net/http"

	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LoginForm struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}

func main() {
	client, err := mongo.Connect(options.Client().
		ApplyURI("[censored for now]"))
	if err != nil {
		panic(err)
	}

	fmt.Println("[Database] Connection Made")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	http.HandleFunc("/login", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		dat, _ := json.Marshal(map[string]interface{}{
			"session": "sedgfoQAI9QAWU9W",
		})
		w.Write(dat)
	}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data *LoginForm = nil
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})
	http.ListenAndServe(":80", nil)
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
