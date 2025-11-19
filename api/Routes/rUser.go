package Routes

import (
	Db "Svelgok-API/Database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	user, session := GetAuthenticatedRequest(c)
	if user == nil || session == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid session",
		})
		return
	}
	user.ID = nil
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func PasswordChange(c *gin.Context) {
	user, session := GetAuthenticatedRequest(c)
	if user == nil || session == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid session",
		})
		return
	}

	var form PasswordChangeForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body received"})
		return
	}

	if form.Original == "" || form.Original == " " || form.New == "" || form.New == " " {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body received"})
		return
	}

	if !user.PasswordMatch(form.Original) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Old password does not match"})
		return
	}

	user, err := user.ChangePassword(form.New)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error changing password " + err.Error()})
		return
	}

	user.InvalidateSessions()

	session, err = Db.Connection.CreateSession(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating new session " + err.Error()})
		return
	}

	claims := user.Claims(session.ID.Hex())

	token, err := Db.CreateJWTToken(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating JWT token " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session": token,
	})
}
