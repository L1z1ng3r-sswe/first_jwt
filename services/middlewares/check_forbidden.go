package middlewares

import (
	"authorisation_app/api/helpers"
	"authorisation_app/db"
	"authorisation_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ForbiddenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token := c.GetHeader("access_token")
		if access_token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user_id, err := helpers.VerifyToken(access_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "second_error": err.Error()})
			c.Abort()
			return
		}

		userId := int(user_id)
		var FoundationProduct models.TProduct
	    param_id:= c.Param("id")

		paramId, err := strconv.Atoi(param_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid product id"})
			c.Abort()
			return
		}

		if err := db.DB.First(&FoundationProduct, paramId).Error ; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found", "second_error": err.Error()})
			return
		}

		if userId != FoundationProduct.UserId {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to delete this product", "matches":   []int{userId, FoundationProduct.UserId}})
			c.Abort()
			return
		}


		c.Set("param_id", paramId)
		c.Next()
	}
}
