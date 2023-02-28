package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthedUser struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Missing `Authorization` header."})
			return
		}

		c.Set("authedUser", "lfkfkfkfk")
		c.Next()
	}
}