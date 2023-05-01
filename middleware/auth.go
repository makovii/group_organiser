package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/makovii/group_organiser/config"
	"strings"
	"github.com/golang-jwt/jwt"
	"fmt"
)

type AuthedUser struct {
	Id             float64  `json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Role  				 float64  `json:"role"`
}

func IsAuthorized(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Missing `Authorization` header."})
			return
		}

		tokenString := strings.Split(header, " ")[1]

		var mySigningKey = []byte(cfg.Secrets.Secret)
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		
    if err != nil {
			fmt.Println("There was an error in parsing", err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Expired token."})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			authedUser := AuthedUser{
				Id:             claims["id"].(float64),
				Name:           claims["name"].(string),
				Email:          claims["email"].(string),
				Role:						claims["role"].(float64),
			}
			c.Set("authedUser", authedUser)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token."})
		}
	}
}