package main

import (
	"authorisation_app/api/routes"
	"authorisation_app/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.New()
	routes.SetUpRoutes(router)

	dbErr := db.InitDB()
	if dbErr != nil {
		log.Fatal("Failed to connect to the database", dbErr)
	}

	defer db.CloseDB()

	if err := router.Run(); err != nil {
		log.Fatal("Failed to launch server")
	}
}
