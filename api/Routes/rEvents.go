package Routes

import (
	Db "Svelgok-API/Database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}

// CreateEvent handles the creation of a new calendar event.
// It expects a JSON body with "title", "description", and "date".
func CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	event := &Db.Event{
		Owner:       user.ID,
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
	}

	res, err := Db.Connection.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetEvents returns all events belonging to the authenticated user.
func GetEvents(c *gin.Context) {
	user, _ := GetAuthenticatedRequest(c)

	events, err := Db.Connection.GetEvents(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

type UpdateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}

// UpdateEvent handles updating an existing event.
// It expects "title", "description", and "date" in the JSON body.
// The event ID is passed as a URL parameter.
func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var req UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.UpdateEvent(id, user.ID, req.Title, req.Description, req.Date); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DeleteEvent handles the deletion of an event.
// The event ID is passed as a URL parameter.
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.DeleteEvent(id, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
