package routes

import (
	"authorisation_app/api/handlers"
	"authorisation_app/services/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {

	// user routes

	router.POST("/signup", handlers.SignUp)
	router.POST("/login", handlers.Login)
	router.POST("/refresh", handlers.Refresh)

	// products routes

	// productRoutes := router.Group("/api/")

	router.GET("/products", handlers.GetAllProducts)
	router.GET("/products/:id", handlers.GetProductByID)
	router.POST("/products",middlewares.AuthMiddleware(), handlers.CreateProduct)
	router.PUT("/products/:id",middlewares.ForbiddenMiddleware(), handlers.EditProduct)
	router.DELETE("/products/:id",middlewares.ForbiddenMiddleware(), handlers.DeleteProduct)
}