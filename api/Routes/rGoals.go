package Routes

import (
	Db "Svelgok-API/Database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGoalRequest struct {
	Title       string       `json:"title" binding:"required"`
	Description string       `json:"description"`
	SubGoals    []Db.SubGoal `json:"subGoals"`
}

// CreateGoal handles the creation of a new goal.
func CreateGoal(c *gin.Context) {
	var req CreateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	goal := &Db.Goal{
		Owner:       user.ID,
		Title:       req.Title,
		Description: req.Description,
		SubGoals:    req.SubGoals,
	}

	res, err := Db.Connection.CreateGoal(goal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetGoals returns all goals belonging to the authenticated user.
func GetGoals(c *gin.Context) {
	user, _ := GetAuthenticatedRequest(c)

	goals, err := Db.Connection.GetGoals(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, goals)
}

// UpdateGoal handles updating an existing goal.
func UpdateGoal(c *gin.Context) {
	id := c.Param("id")
	var req UpdateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.UpdateGoal(id, user.ID, req.Title, req.Description, req.SubGoals); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DeleteGoal handles the deletion of a goal.
func DeleteGoal(c *gin.Context) {
	id := c.Param("id")
	user, _ := GetAuthenticatedRequest(c)

	if err := Db.Connection.DeleteGoal(id, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
