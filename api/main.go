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

type NotesForm struct {
	Notes   string `json:"notes"`
	Session string `json:"session"`
}

type SessionAuth struct {
	Session string `json:"s"`
}

// ==========================
// Exercise data model
// ==========================
type Exercise struct {
	Muscle    string   `json:"muscle"`
	Exercises []string `json:"exercises"`
}

// Handler for /exercise
func exerciseHandler(w http.ResponseWriter, r *http.Request) {
	// Get query param ?muscle=chest
	muscle := r.URL.Query().Get("muscle")

	// Predefined exercise list
	exerciseMap := map[string][]string{
		"chest":     {"Push-ups", "Bench Press", "Chest Fly"},
		"back":      {"Pull-ups", "Deadlift", "Barbell Row"},
		"legs":      {"Squats", "Lunges", "Leg Press"},
		"arms":      {"Bicep Curls", "Tricep Dips", "Hammer Curls"},
		"shoulders": {"Overhead Press", "Lateral Raises", "Arnold Press"},
	}

	// Default response if no muscle is found
	exercises, ok := exerciseMap[muscle]
	if !ok {
		http.Error(w, "Muscle not found. Try chest, back, legs, arms, shoulders.", http.StatusBadRequest)
		return
	}

	// Build response
	response := Exercise{
		Muscle:    muscle,
		Exercises: exercises,
	}

	// Send JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Load environment variables
	Config.Load()
	// Connect Database
	_, err := Database.NewDatabase(Config.DATABASE_CONNECTION_STRING)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Database] Connection Made")

	// Disconnect the database when main ends
	defer func() {
		if err := Database.Connection.Disconnect(); err != nil {
			panic(err)
		}
	}()

	// LOGIN route
	http.HandleFunc("/login", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		var LoginReq *LoginForm = &LoginForm{}
		err := Util.DeserializeRequest(&LoginReq, r)
		if err != nil {
			return
		}

		fmt.Println(LoginReq)

		user, found, err := Database.Connection.UserFindUsername(LoginReq.Username)
		if err != nil {
			return
		}
		if !found {
			return
		}

		if Util.Hash256(LoginReq.Password) != user.Password {
			return
		}

		session, err := Database.Connection.SessionCreate(user, Util.RandomString(12))
		if err != nil {
			return
		}

		Util.SendJSON(200, map[string]interface{}{
			"session": session.SessionID,
		}, w)
	}))

	// NOTES route
	http.HandleFunc("/notes", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		var NotesReq *NotesForm = &NotesForm{}
		err := Util.DeserializeRequest(&NotesReq, r)
		if err != nil {
			return
		}

		session, found, err := Database.Connection.SessionFind(NotesReq.Session)
		if err != nil {
			return
		}

		if !found {
			return
		}

		user, found, err := Database.Connection.UserFindUsername(session.Target)
		if err != nil {
			return
		}
		if !found {
			return
		}

		user.Password = ""

		user, _, err = Database.Connection.UpdateNotes(user.Username, NotesReq.Notes)

		if err != nil {
			// Error updating the notes
		}

		user.Password = ""

		Util.SendJSON(200, user, w)

	}))

	// SESSION route
	http.HandleFunc("/session", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		var SessionReq *SessionAuth = &SessionAuth{}
		err := Util.DeserializeRequest(&SessionReq, r)
		if err != nil {
			return
		}

		session, found, err := Database.Connection.SessionFind(SessionReq.Session)
		if err != nil {
			return
		}

		if !found {
			return
		}

		user, found, err := Database.Connection.UserFindUsername(session.Target)
		if err != nil {
			return
		}
		if !found {
			return
		}

		user.Password = ""
		Util.SendJSON(200, user, w)

	}))

	// ==========================
	// EXERCISE route
	// ==========================
	http.HandleFunc("/exercise", corsMiddleware(exerciseHandler))

	// ROOT route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data *LoginForm = nil
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	// Start server
	http.ListenAndServe(":7284", nil)
}

// ==========================
// CORS Middleware
// ==========================
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
