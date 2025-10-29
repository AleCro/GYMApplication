// Alejandro Coro Lopez
// 202610 Open Source Web Technologies CEN-4350C-14149
// October 7, 2025
//
// Full GYMApp API with Notes (array support), Calendar integration,
// Progress tracking, and Goal management (MongoDB-backed).

package main

import (
	"GYMAppAPI/Config"
	"GYMAppAPI/Database"
	"GYMAppAPI/Util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

// Update Note form structure
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

// Progress model
type ProgressEntry struct {
	Date     string  `json:"date"`
	Weight   float64 `json:"weight"`
	Message  string  `json:"message"`
	PhotoURL string  `json:"photo"`
}

// Goal model (simple for MongoDB)
type Goal struct {
	Title     string   `json:"title"`
	Steps     []string `json:"steps"`
	Completed bool     `json:"completed"`
}

// ==========================
// IN-MEMORY PLACEHOLDER (Sessions Map)
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

// ==========================
// HANDLERS
// ==========================

// LOGIN route
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var LoginReq *LoginForm = &LoginForm{}
	err := Util.DeserializeRequest(&LoginReq, r)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFindUsername(LoginReq.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if Util.Hash256(LoginReq.Password) != user.Password {
		http.Error(w, "Passwords do not match", http.StatusUnauthorized)
		return
	}

	session, err := Database.Connection.SessionCreate(user, Util.RandomString(12))
	if err != nil {
		http.Error(w, "Internal Server Error while creating Session", http.StatusInternalServerError)
		return
	}

	sendJSON(200, map[string]any{"session": session.SessionID}, w)
}

// ADD NOTE route
func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req AddNoteForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	Database.Connection.AddNoteToUser(user.Username, req.Note)
	sendJSON(200, map[string]any{"success": true}, w)
}

// ADD EVENT route
func addEventHandler(w http.ResponseWriter, r *http.Request) {
	var req AddEvent
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	Database.Connection.AddCalendarEvent(user.Username, &Database.CalendarEvent{
		Name: req.Title,
		Time: req.Time,
	})

	sendJSON(200, map[string]any{"success": true}, w)
}

// UPDATE NOTE route
func updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateNoteForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	err = Database.Connection.UpdateNote(user.Username, req.Index, req.Note)
	if err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
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
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	err = Database.Connection.DeleteNote(user.Username, req.Index)
	if err != nil {
		http.Error(w, "Could not delete note", http.StatusInternalServerError)
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

	sendJSON(200, map[string]any{"muscle": muscle, "exercises": exercises}, w)
}

// ROOT route
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("GYMApp API running ✅"))
}

// SESSION route
func sessionHandler(w http.ResponseWriter, r *http.Request) {
	var req SessionAuth
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	session, found, err := Database.Connection.SessionFind(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	user, found, err := Database.Connection.UserFindUsername(session.Target)
	if err != nil || !found {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	user.Password = ""
	sendJSON(200, user, w)
}

// ==========================
// PROGRESS & GOAL HANDLERS (MongoDB-backed)
// ==========================

type AddProgressForm struct {
	Session string  `json:"session"`
	Weight  float64 `json:"weight"`
	Message string  `json:"message"`
	Photo   string  `json:"photo"`
}

type AddGoalForm struct {
	Session string   `json:"session"`
	Title   string   `json:"title"`
	Steps   []string `json:"steps"`
}

type UpdateGoalForm struct {
	Session string `json:"session"`
	Index   int    `json:"i"`
	Done    bool   `json:"done"`
}

// ADD PROGRESS
func addProgressHandler(w http.ResponseWriter, r *http.Request) {
	var req AddProgressForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	entry := Database.ProgressEntry{
		Date:     Util.CurrentDate(),
		Weight:   req.Weight,
		Message:  req.Message,
		PhotoURL: req.Photo,
	}

	err = Database.Connection.AddProgress(user.Username, entry)
	if err != nil {
		http.Error(w, "Could not save progress", http.StatusInternalServerError)
		return
	}

	sendJSON(200, map[string]any{"success": true, "progress": entry}, w)
}

// GET PROGRESS
func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	var req SessionAuth
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	progress, err := Database.Connection.GetProgress(user.Username)
	if err != nil {
		http.Error(w, "Could not load progress", http.StatusInternalServerError)
		return
	}

	sendJSON(200, progress, w)
}

// ADD GOAL
func addGoalHandler(w http.ResponseWriter, r *http.Request) {
	var req AddGoalForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	goal := Database.Goal{
		Title:     req.Title,
		Steps:     req.Steps,
		Completed: false,
	}

	err = Database.Connection.AddGoal(user.Username, goal)
	if err != nil {
		http.Error(w, "Could not save goal", http.StatusInternalServerError)
		return
	}

	sendJSON(200, map[string]any{"success": true, "goal": goal}, w)
}

// GET GOALS
func getGoalsHandler(w http.ResponseWriter, r *http.Request) {
	var req SessionAuth
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	goals, err := Database.Connection.GetGoals(user.Username)
	if err != nil {
		http.Error(w, "Could not load goals", http.StatusInternalServerError)
		return
	}

	sendJSON(200, goals, w)
}

// UPDATE GOAL
func updateGoalHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateGoalForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	err = Database.Connection.UpdateGoal(user.Username, req.Index, req.Done)
	if err != nil {
		http.Error(w, "Could not update goal", http.StatusInternalServerError)
		return
	}

	sendJSON(200, map[string]any{"success": true}, w)
}

// ==========================
// UPLOAD HANDLER
// ==========================
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(20 << 20) // limit 20MB

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Missing image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// create uploads folder if missing
	os.MkdirAll("uploads", os.ModePerm)

	// generate unique name
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%s_%s%s", Util.CurrentDate(), Util.RandomString(8), ext)
	path := filepath.Join("uploads", filename)

	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	Util.SendJSON(http.StatusOK, map[string]any{
		"url":     fmt.Sprintf("/uploads/%s", filename),
		"success": true,
	}, w)
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
	http.HandleFunc("/exercise", corsMiddleware(exerciseHandler))

	// Progress & Goals endpoints
	http.HandleFunc("/addprogress", corsMiddleware(addProgressHandler))
	http.HandleFunc("/getprogress", corsMiddleware(getProgressHandler))
	http.HandleFunc("/addgoal", corsMiddleware(addGoalHandler))
	http.HandleFunc("/getgoals", corsMiddleware(getGoalsHandler))
	http.HandleFunc("/updategoal", corsMiddleware(updateGoalHandler))

	http.HandleFunc("/upload", corsMiddleware(uploadHandler))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	http.ListenAndServe(":7284", nil)
}
