package handlers

import (
	"authorisation_app/api/helpers"
	"authorisation_app/db"
	"log"

	"authorisation_app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var newUserWithHashedPassword models.TUser
	var newUser models.TUser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request to create user"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "second_error": err.Error()})
		return
	}

	hashedPassword, err := helpers.PasswordHasher(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to hash password: %v", err)})
		return
	}

	newUserWithHashedPassword = models.TUser{
		Email:    newUser.Email,
		Password: hashedPassword,
	}

	err = db.DB.Create(&newUserWithHashedPassword).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is already exist"})
		return
	}

	accessToken ,err := helpers.GenerateAccessToken(int(newUserWithHashedPassword.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken ,err := helpers.GenerateRefreshToken(int(newUserWithHashedPassword.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	response := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context){
	var loginUser models.TUser
	err := c.ShouldBindJSON(&loginUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials?"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed"})
		return
	}

	var userFoundation models.TUser
	if err := db.DB.Where("email = ?", loginUser.Email).First(&userFoundation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFoundation.Password),  []byte(loginUser.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"} )
		return
	}

	accessToken ,err := helpers.GenerateAccessToken(int(userFoundation.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}
	
	refreshToken ,err := helpers.GenerateRefreshToken(int(userFoundation.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	response := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusOK, response)
} 

func Refresh(c *gin.Context){
	refreshToken := c.GetHeader("refresh_token")
	log.Println(refreshToken)

	if refreshToken == " "{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get refresh token"})
		return
	}

	user_id ,err := helpers.VerifyToken(refreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token", "second_error": err.Error()})
		return
	}

	newAccessToken, err := helpers.GenerateAccessToken(int(user_id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access-token", "second_error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}