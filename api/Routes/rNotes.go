package Routes

import (
	Db "Svelgok-API/Database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateNote handles the creation of a new note.
// It expects a JSON body with "title" and "content".
func CreateNote(c *gin.Context) {
	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	note := &Db.Note{
		Owner:   user.ID,
		Title:   req.Title,
		Content: req.Content,
	}

	res, err := Db.Connection.CreateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetNotes returns all notes belonging to the authenticated user.
func GetNotes(c *gin.Context) {
	user, _ := GetAuthenticatedRequest(c)

	notes, err := Db.Connection.GetNotes(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

type UpdateNoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateNote handles updating an existing note.
// It expects "title" and "content" in the JSON body.
// The note ID is passed as a URL parameter.
func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var req UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.UpdateNote(id, user.ID, req.Title, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DeleteNote handles the deletion of a note.
// The note ID is passed as a URL parameter.
func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.DeleteNote(id, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
