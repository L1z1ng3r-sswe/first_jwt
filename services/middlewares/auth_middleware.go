package middlewares

import (
	"authorisation_app/api/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token := c.GetHeader("access_token")
		log.Println("Received Token:", access_token)
		if access_token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user_id, err := helpers.VerifyToken(access_token)
		if err != nil {
			log.Println("Token Verification Failed:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "second_error": err.Error()})
			c.Abort()
			return
		}

		log.Println("User ID from Token:", user_id)

		c.Set("user_id", int(user_id))
		c.Next()
	}
}
