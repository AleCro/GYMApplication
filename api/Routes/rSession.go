package Routes

import (
	Db "Svelgok-API/Database"
	"Svelgok-API/Environment"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateJWTSession(c *gin.Context) {
	var form LoginRequestForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body received"})
		return
	}
	user, found, err := Db.Connection.FilterOneUser(bson.D{{"username", form.Username}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error while fetching users " + err.Error(),
		})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No user with that username was found",
		})
		return
	}
	if !user.PasswordMatch(form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Passwords do not match",
		})
		return
	}

	session, err := Db.Connection.CreateSession(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error creating session " + err.Error(),
		})
		return
	}

	// JWT token creation
	token, err := Db.CreateJWTToken(user.Claims(session.ID.Hex()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error while creating JWT token: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session": token,
	})
}

func RefreshJWTSession(c *gin.Context) {
	claims := GetCheckedClaims(c)
	user, session := GetAuthenticatedRequest(c)
	if user == nil || session == nil || claims == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid session",
		})
		return
	}

	var new_expiration time.Time = time.Now().Add(Environment.JWT_TOKEN_LIFESPAN)

	if new_expiration.After(session.CreatedAt.Add(Environment.SESSION_DURATION)) {
		new_expiration = session.CreatedAt
	}

	now := time.Now()

	if now.After(claims.ExpiresAt.Time.Add(Environment.JWT_RENEWAL_GRACE_PERIOD)) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to renew token",
		})
		return
	}

	if Environment.STRICT_SESSION_CONSISTENCY {
		// Manual session expiration check
		if now.After(session.CreatedAt.Add(Environment.SESSION_DURATION + Environment.STRICT_SESSION_CONSISTENCY_OFFSET)) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to renew token",
			})
			return
		}

		// Check JWT lifespan is consistent with environment variables
		if claims.IssuedAt.Add(Environment.JWT_TOKEN_LIFESPAN+Environment.JWT_RENEWAL_GRACE_PERIOD).Sub(claims.ExpiresAt.Time).Abs() > Environment.STRICT_SESSION_CONSISTENCY_OFFSET {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to renew token",
			})
			return
		}
	}

	token, err := Db.CreateJWTToken(user.Claims(session.ID.Hex()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parsing JWT token " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"session": token,
	})
}

// Removes the session from the database, making any JWT token
// not pass `RequireSessionValidate` neccesary to make any changes.
func RemoveSession(c *gin.Context) {
	session := GetCheckedClaims(c)
	if session == nil {
		c.JSON(http.StatusNotModified, gin.H{})
		return
	}

	deleted, err := Db.Connection.RemoveSessionFromID(session.SessionID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error " + err.Error(),
		})
		return
	}

	if deleted == 0 {
		c.JSON(http.StatusNotModified, gin.H{
			"message": "No session to delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
