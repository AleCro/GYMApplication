package main

import (
	"GYMAppAPI/Config"
	"GYMAppAPI/Database"
	"GYMAppAPI/Util"
	"fmt"
	"net/http"

	"encoding/json"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionAuth struct {
	Session string `json:"s"`
}

func main() {
	// Load environment vairables
	Config.Load()
	// Connect Database
	_, err := Database.NewDatabase(Config.DATABASE_CONNECTION_STRING)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Database] Connection Made")

	// This will disconnect the database at the end of the main function
	defer func() {
		if err := Database.Connection.Disconnect(); err != nil {
			panic(err)
		}
	}()

	http.HandleFunc("/login", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		var LoginReq *LoginForm = &LoginForm{}
		err := Util.DeserializeRequest(&LoginReq, r)
		if err != nil {
			return
		}

		user, found, err := Database.Connection.UserFindUsername(LoginReq.Username)
		if err != nil {
			// Internal Server Error
			return
		}
		if !found {
			// User not found
			return
		}

		if Util.Hash256(LoginReq.Password) != user.Password {
			// Passwords do NOT match
			return
		}

		session, err := Database.Connection.SessionCreate(user, Util.RandomString(12))
		if err != nil {
			// Internal Server Error
			return
		}

		Util.SendJSON(200, map[string]interface{}{
			"session": session.SessionID,
		}, w)
		return
	}))

	http.HandleFunc("/session", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		var SessionReq *SessionAuth = &SessionAuth{}
		err := Util.DeserializeRequest(&SessionReq, r)
		if err != nil {
			return
		}

		session, found, err := Database.Connection.SessionFind(SessionReq.Session)
		if err != nil {
			// Internal Server Error
			return
		}

		if !found {
			// Session not found
			return
		}

		user, found, err := Database.Connection.UserFindUsername(session.Target)
		if err != nil {
			// Internal Server Error
			return
		}
		if !found {
			// User not found
			return
		}

		user.Password = ""
		Util.SendJSON(200, user, w)

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
