// Alejandro Coro Lopez
// 202610 Open Source Web Technologies CEN-4350C-14149
// October 7, 2025
//
// Full GYMApp API with Notes (array support) and Calendar integration.
// This version uses in-memory storage for demonstration purposes.
// You can later connect this to your real Database package.

package main

import (
	"GYMAppAPI/Config"
	"GYMAppAPI/Database"
	"GYMAppAPI/Util"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// ==========================
// STRUCTS & DATA MODELS
// ==========================

// Login form structure
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Add Note form structure
type AddNoteForm struct {
	Session string `json:"session"`
	Note    string `json:"note"`
}

// Add Event form structure
type AddEvent struct {
	Session string `json:"session"`
	Title   string `json:"title"`
	Time    uint64 `json:"time"`
}

// Add Note form structure
type UpdateNoteForm struct {
	Session string `json:"session"`
	Index   int    `json:"i"`
	Note    string `json:"note"`
}

// Delete Note form structure
type DeleteNoteForm struct {
	Session string `json:"session"`
	Index   int    `json:"i"`
}

// Session request structure
type SessionAuth struct {
	Session string `json:"s"`
}

// Exercise model
type Exercise struct {
	Muscle    string   `json:"muscle"`
	Exercises []string `json:"exercises"`
}

// ==========================
// IN-MEMORY DATABASE (DEMO)
// ==========================

var (
	users    = make(map[string]*Database.Database)
	sessions = make(map[string]string)
	lock     sync.Mutex
)

// ==========================
// UTILITY FUNCTIONS
// ==========================

func sendJSON(status int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[i%len(letters)]
	}
	return string(s)
}

// ==========================
// HANDLERS
// ==========================

// LOGIN route
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var LoginReq *LoginForm = &LoginForm{}
	err := Util.DeserializeRequest(&LoginReq, r)
	if err != nil {
		return
	}


	user, found, err := Database.Connection.UserFindUsername(LoginReq.Username)
	if err != nil {
		http.Error(w, "{\"error\": \"Internal Server Error\"}", http.StatusInternalServerError)
		return
	}
	if !found {
		http.Error(w, "{\"error\": \"User not Found\"}", http.StatusNotFound)
		return
	}

	if Util.Hash256(LoginReq.Password) != user.Password {
		http.Error(w, "{\"error\": \"Passwords do not match\"}", http.StatusUnauthorized)
		return
	}

	session, err := Database.Connection.SessionCreate(user, Util.RandomString(12))
	if err != nil {
		http.Error(w, "{\"error\": \"Internal Server Error while creating Session\"}", http.StatusInternalServerError)
		return
	}

	Util.SendJSON(200, map[string]interface{}{
		"session": session.SessionID,
	}, w)
}

// ADD NOTE route
func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req AddNoteForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if !found {
		return
	}

	if err != nil {
		return
	}

	user.Notes = append(user.Notes, req.Note)

	Database.Connection.AddNoteToUser(user.Username, req.Note)

	sendJSON(200, user, w)
}

// ADD EVENT route
func addEventHandler(w http.ResponseWriter, r *http.Request) {
	var req AddEvent
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if !found {
		return
	}

	if err != nil {
		return
	}

	Database.Connection.AddCalendarEvent(user.Username, &Database.CalendarEvent{
		Name: req.Title,
		Time: req.Time,
	})

	sendJSON(200, user, w)
}

// UPDATE NOTE route
func updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateNoteForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if !found {
		return
	}

	if err != nil {
		return
	}

	if req.Index > len(user.Notes)-1 {
		return
	}

	err = Database.Connection.UpdateNote(user.Username, req.Index, req.Note)

	if err != nil {
		return
	}

	sendJSON(200, map[string]any{"success": true}, w)
}

// DELETE NOTE route
func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteNoteForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if !found {
		return
	}

	if err != nil {
		return
	}

	if req.Index > len(user.Notes)-1 {
		return
	}

	err = Database.Connection.DeleteNote(user.Username, req.Index)

	if err != nil {
		return
	}

	sendJSON(200, map[string]any{"success": true}, w)
}

// EXERCISE route
func exerciseHandler(w http.ResponseWriter, r *http.Request) {
	muscle := r.URL.Query().Get("muscle")

	exercises, ok := ExerciseMap[muscle]
	if !ok {
		http.Error(w, "Muscle not found", http.StatusBadRequest)
		return
	}

	response := map[string]any{
		"muscle":    muscle,
		"exercises": exercises,
	}
	sendJSON(200, response, w)
}

// ROOT route
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("GYMApp API running ✅"))
}

// Handle the fetching of sessions
func sessionHandler(w http.ResponseWriter, r *http.Request) {
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

}

// ==========================
// MIDDLEWARE
// ==========================
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// ==========================
// MAIN
// ==========================
func main() {
	// Load environment variables
	Config.Load()
	// Connect Database
	_, err := Database.NewDatabase(Config.DATABASE_CONNECTION_STRING)
	// hello

	if err != nil {
		panic(err)
	}

	fmt.Println("[Server] GYMApp API starting on port 7284...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/session", corsMiddleware(sessionHandler))
	http.HandleFunc("/login", corsMiddleware(loginHandler))
	http.HandleFunc("/addnote", corsMiddleware(addNoteHandler))
	http.HandleFunc("/addevent", corsMiddleware(addEventHandler))
	http.HandleFunc("/updatenote", corsMiddleware(updateNoteHandler))
	http.HandleFunc("/deletenote", corsMiddleware(deleteNoteHandler))
	http.HandleFunc("/calendar", corsMiddleware(addEventHandler))
	http.HandleFunc("/exercise", corsMiddleware(exerciseHandler))

	http.ListenAndServe(":7284", nil)
}
