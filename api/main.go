package main

import (
	"GYMAppAPI/Config"
	"GYMAppAPI/Database"
	"GYMAppAPI/Exercise"
	"GYMAppAPI/Util"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	// Use google/uuid for generating unique IDs
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ==========================
// GLOBAL CONSTANTS
// ==========================
const UploadDir = "./uploads"

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
	Session  string `json:"session"`
	Title    string `json:"title"`
	Time     uint64 `json:"time"`
	Timezone string `json:"timezone,omitempty"`
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

// Generic delete form (shared by goals & progress)
type DeleteForm struct {
	Session string `json:"session"`
	Index   int    `json:"i"`
}

// Session request structure
type SessionAuth struct {
	Session string `json:"session"`
}

type AddGoalForm struct {
	Session string   `json:"session"`
	Title   string   `json:"title"`
	Steps   []string `json:"steps"`
}

type DeleteGoalForm struct {
	Session string `json:"session"`
	Index   int    `json:"i"`
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
	if err != nil {
		return
	}
	if !found {
		return
	}
	user.Notes = append(user.Notes, req.Note)

	Database.Connection.AddNoteToUser(user.Username, req.Note)
	user.Password = ""
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
		Name:     req.Title,
		Time:     req.Time,
		Timezone: req.Timezone,
	})

	sendJSON(200, user, w)
}

// DELETE EVENT route
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DEBUG] /deleteevent called")

	var req struct {
		Session string `json:"session"`
		Index   int    `json:"i"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("[DEBUG] Decode error:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	fmt.Printf("[DEBUG] Session: %s, Index: %d\n", req.Session, req.Index)

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil {
		fmt.Println("[DEBUG] UserFromSession error:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if !found {
		fmt.Println("[DEBUG] Session not found")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := Database.Connection.DeleteCalendarEvent(user.Username, req.Index); err != nil {
		fmt.Println("[DEBUG] DeleteCalendarEvent error:", err)
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	fmt.Println("[DEBUG] Event deleted successfully for user:", user.Username)
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

	exercises, ok := Exercise.ExerciseMap[muscle]
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
// UPLOAD HANDLER (Base64-Compatible, No Filesystem Needed)
// ==========================
func uploadProgressHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🟢 [uploadProgressHandler] triggered")

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Println("❌ ParseMultipartForm error:", err)
		sendJSON(400, map[string]any{
			"success": false,
			"message": "Invalid upload form",
		}, w)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("❌ FormFile error:", err)
		sendJSON(400, map[string]any{
			"success": false,
			"message": "No image file provided",
		}, w)
		return
	}
	defer file.Close()

	// Convert file to Base64 directly (no saving)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("❌ Read file error:", err)
		sendJSON(500, map[string]any{
			"success": false,
			"message": "Error reading file",
		}, w)
		return
	}

	mimeType := http.DetectContentType(fileBytes)
	base64Str := fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(fileBytes))

	fmt.Printf("✅ Uploaded %s as Base64 (%d bytes)\n", handler.Filename, len(fileBytes))

	sendJSON(200, map[string]any{
		"success": true,
		"message": "Image converted to Base64 successfully",
		"base64":  base64Str,
	}, w)
}

// ==========================
// ADD PROGRESS (Return entry with string ID)
// ==========================
func addProgressHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Session string  `json:"session"`
		Weight  float64 `json:"weight"`
		Message string  `json:"message"`
		Photo   string  `json:"photo"` // base64 or URL
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSON(400, map[string]any{"success": false, "message": "Invalid JSON body"}, w)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		sendJSON(401, map[string]any{"success": false, "message": "Invalid session"}, w)
		return
	}

	entry := Database.ProgressEntry{
		ID:      primitive.NewObjectID(),
		Date:    time.Now().Format("2006-01-02"),
		Weight:  req.Weight,
		Message: req.Message,
		Photo:   req.Photo,
	}
	if err := Database.Connection.AddProgress(user.Username, entry); err != nil {
		sendJSON(500, map[string]any{"success": false, "message": fmt.Sprintf("Failed to save progress: %v", err)}, w)
		return
	}

	sendJSON(200, map[string]any{
		"success": true,
		"message": "Progress saved successfully!",
		"data": map[string]any{
			"id":          entry.ID.Hex(),
			"date":        entry.Date,
			"weight":      entry.Weight,
			"message":     entry.Message,
			"photoBase64": entry.Photo,
		},
	}, w)
}

// ==========================
// GET PROGRESS (Normalize IDs for frontend)
// ==========================
func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Session string `json:"session"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSON(400, map[string]any{"success": false, "message": "Invalid request"}, w)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		sendJSON(401, map[string]any{"success": false, "message": "Invalid session"}, w)
		return
	}

	progress, err := Database.Connection.GetProgress(user.Username)
	if err != nil {
		sendJSON(500, map[string]any{"success": false, "message": "Failed to fetch progress"}, w)
		return
	}

	out := make([]map[string]any, 0, len(progress))
	for _, p := range progress {
		out = append(out, map[string]any{
			"id":          p.ID.Hex(),
			"date":        p.Date,
			"weight":      p.Weight,
			"message":     p.Message,
			"photoBase64": p.Photo,
		})
	}
	sendJSON(200, out, w)
}

// ==========================
// DELETE PROGRESS
// ==========================
// DELETE PROGRESS HANDLER
func deleteProgressHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Session string `json:"session"`
		ID      string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		http.Error(w, "Invalid progress ID", http.StatusBadRequest)
		return
	}

	if err := Database.Connection.DeleteProgress(user.Username, objID); err != nil {
		http.Error(w, fmt.Sprintf("Delete failed: %v", err), http.StatusInternalServerError)
		return
	}
	sendJSON(200, map[string]any{"success": true, "message": "Progress entry deleted!"}, w)
}

// ==========================
// ADD GOAL
// ==========================
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

// ==========================
// GET GOALS
// ==========================
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

// ==========================
// DELETE GOAL - NEW HANDLER
// ==========================
func deleteGoalHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteGoalForm
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	// Assuming Database.Connection has a DeleteGoal method
	err = Database.Connection.DeleteGoal(user.Username, req.Index)
	if err != nil {
		http.Error(w, "Could not delete goal", http.StatusInternalServerError)
		return
	}

	sendJSON(200, map[string]any{"success": true}, w)
}

// ==========================
// UPDATE GOAL (Handles completion + steps)
// ==========================
func updateGoalHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Session string   `json:"session"`
		Index   int      `json:"i"`
		Done    *bool    `json:"done,omitempty"`
		Steps   []string `json:"steps,omitempty"`
		NewStep bool     `json:"newStep,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, found, err := Database.Connection.UserFromSession(req.Session)
	if err != nil || !found {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	//  If "Steps" is provided → update steps
	if req.Steps != nil {
		if err := Database.Connection.UpdateGoalSteps(user.Username, req.Index, req.Steps); err != nil {
			http.Error(w, "Could not update goal steps", http.StatusInternalServerError)
			return
		}
		sendJSON(200, map[string]any{"success": true, "updated": "steps"}, w)
		return
	}

	//  Otherwise, treat as a completion update
	if req.Done != nil {
		if err := Database.Connection.UpdateGoal(user.Username, req.Index, *req.Done); err != nil {
			http.Error(w, "Could not update goal", http.StatusInternalServerError)
			return
		}
		sendJSON(200, map[string]any{"success": true, "updated": "completion"}, w)
		return
	}

	http.Error(w, "No valid fields provided", http.StatusBadRequest)
}

// ID PATCH
func patchProgressIDsHandler(w http.ResponseWriter, r *http.Request) {
	coll := Database.Connection.Client().Database(Config.DATABASE_NAME).Collection(Config.DATABASE_USER_COLLECTION)

	cursor, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer cursor.Close(context.Background())

	patchedUsers := 0
	for cursor.Next(context.Background()) {
		var user Database.User
		if err := cursor.Decode(&user); err != nil {
			continue
		}
		updated := false
		for i := range user.Progress {
			if user.Progress[i].ID.IsZero() {
				user.Progress[i].ID = primitive.NewObjectID()
				updated = true
			}
		}
		if updated {
			_, err := coll.UpdateOne(
				context.Background(),
				bson.M{"username": user.Username},
				bson.M{"$set": bson.M{"progress": user.Progress}},
			)
			if err == nil {
				patchedUsers++
			}
		}
	}
	sendJSON(200, map[string]any{"patched_users": patchedUsers}, w)
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

	http.HandleFunc("/session", corsMiddleware(sessionHandler))
	http.HandleFunc("/login", corsMiddleware(loginHandler))
	http.HandleFunc("/addnote", corsMiddleware(addNoteHandler))
	http.HandleFunc("/addevent", corsMiddleware(addEventHandler))
	http.HandleFunc("/updatenote", corsMiddleware(updateNoteHandler))
	http.HandleFunc("/deletenote", corsMiddleware(deleteNoteHandler))
	http.HandleFunc("/calendar", corsMiddleware(addEventHandler))
	http.HandleFunc("/deleteevent", corsMiddleware(deleteEventHandler))
	http.HandleFunc("/exercise", corsMiddleware(exerciseHandler))

	// Progress & Goals endpoints
	http.HandleFunc("/addprogress", corsMiddleware(addProgressHandler))
	http.HandleFunc("/getprogress", corsMiddleware(getProgressHandler))
	http.HandleFunc("/deleteprogress", corsMiddleware(deleteProgressHandler))
	http.HandleFunc("/upload", corsMiddleware(uploadProgressHandler))
	http.HandleFunc("/patchprogressids", patchProgressIDsHandler)

	http.HandleFunc("/addgoal", corsMiddleware(addGoalHandler))
	http.HandleFunc("/getgoals", corsMiddleware(getGoalsHandler))
	http.HandleFunc("/updategoal", corsMiddleware(updateGoalHandler))
	http.HandleFunc("/deletegoal", corsMiddleware(deleteGoalHandler))

	// Serve uploaded images
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(UploadDir))))

	// Put this at the very end!
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":7284", nil)
}
