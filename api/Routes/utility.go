package Routes

import (
	Db "Svelgok-API/Database"

	"github.com/gin-gonic/gin"
)

func GetAuthenticatedRequest(c *gin.Context) (*Db.User, *Db.Session) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, nil
	}

	sessionInterface, exists := c.Get("session")
	if !exists {
		return nil, nil
	}

	user, casted := userInterface.(*Db.User)
	if !casted {
		return nil, nil
	}

	session, casted := sessionInterface.(*Db.Session)
	if !casted {
		return user, nil
	}

	return user, session
}

func GetCheckedClaims(c *gin.Context) *Db.UserJWTClaim {
	v, exist := c.Get("claims")
	if !exist {
		return nil
	}
	return v.(*Db.UserJWTClaim)
}

func OPTIONS(c *gin.Context) {
	c.AbortWithStatus(204)
}
