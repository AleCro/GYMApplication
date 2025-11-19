package Routes

import (
	Db "Svelgok-API/Database"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProgress handles the creation of a new progress entry.
// It expects a multipart form with "title", "description", and an "image" file.
// The image is converted to a Base64 string for storage.
func CreateProgress(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	file, err := c.FormFile("image")

	if title == "" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and image are required"})
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
		return
	}
	defer src.Close()

	// Read the file content
	fileBytes := make([]byte, file.Size)
	if _, err := src.Read(fileBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	// Encode to Base64
	base64Image := base64.StdEncoding.EncodeToString(fileBytes)
	// Determine mime type (simple check or assume from extension, for now we can prefix generic or try to detect)
	// For simplicity in display, we might want to store the full data URI scheme or just the base64.
	// Let's store the full data URI if possible, or just base64 and let frontend handle it.
	// The plan said "Base64 display", usually `data:image/xyz;base64,...`
	// Let's construct a simple data URI.
	mimeType := file.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = "image/jpeg" // Fallback
	}
	imageData := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Image)

	user, _ := GetAuthenticatedRequest(c)

	progress := &Db.Progress{
		Owner:       user.ID,
		Title:       title,
		Description: description,
		ImageData:   imageData,
	}

	res, err := Db.Connection.CreateProgress(progress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProgress returns all progress entries for the authenticated user.
func GetProgress(c *gin.Context) {
	user, _ := GetAuthenticatedRequest(c)

	progress, err := Db.Connection.GetProgress(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

// UpdateProgress handles updating the text details of a progress entry.
// It expects "title" and "description" in the JSON body.
// The entry ID is passed as a URL parameter.
func UpdateProgress(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.UpdateProgress(id, user.ID, body.Title, body.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Progress updated"})
}

// DeleteProgress handles the deletion of a progress entry.
// The entry ID is passed as a URL parameter.
func DeleteProgress(c *gin.Context) {
	id := c.Param("id")
	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.DeleteProgress(id, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Progress deleted"})
}
