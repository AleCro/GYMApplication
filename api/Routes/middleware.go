package Routes

import (
	Db "Svelgok-API/Database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SoftSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.TrimSpace(c.GetHeader("Authorization"))
		if auth == "" {
			c.Next()
			return
		}

		var bearer = "Bearer "
		if !strings.HasPrefix(strings.ToLower(auth), strings.ToLower(bearer)) {
			c.Next()
			return
		}

		token := strings.TrimSpace(auth[len(bearer):])
		claims, err := Db.VerifyJWTSignature(token, true)
		if err == nil {
			c.Set("claims", claims)
		}
		c.Next()
	}
}

func RequireSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := GetCheckedClaims(c)
		if claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "authorization required"})
			return
		}

		if claims.Expired() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Expired token",
			})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func RequireSessionValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := GetCheckedClaims(c)
		if claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "authorization required"})
			return
		}

		if claims.Expired() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Expired token",
			})
			return
		}

		user, session, err := Db.ValidateJWTSessionFromClaims(claims, false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired session"})
			return
		}

		c.Set("user", user)
		c.Set("session", session)
		c.Next()
	}
}

func RequireSessionValidateEx() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := GetCheckedClaims(c)
		if claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "authorization required"})
			return
		}

		user, session, err := Db.ValidateJWTSessionFromClaims(claims, true)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired session"})
			return
		}

		c.Set("user", user)
		c.Set("session", session)
		c.Next()
	}
}
