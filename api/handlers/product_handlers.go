package handlers

import (
	"authorisation_app/db"
	"authorisation_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllProducts(c *gin.Context) {
	var allProducts []models.TProduct 

	result := db.DB.Find(&allProducts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products", "second_error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, allProducts)
}

func CreateProduct(c *gin.Context) {
    var newProduct models.TProduct
    userId := c.GetInt("user_id")
	
    if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "second_error": err.Error()})
        return
    }

	validate := validator.New()

	if err := validate.Struct(newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body", "second_error": err.Error()})
		return
	}
	
	newProduct.UserId = userId

    if err := db.DB.Create(&newProduct).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid converter to sql", "second_error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "New product successfully created", "product": newProduct})
}

func GetProductByID(c *gin.Context){
	param_id:=c.Param("id")
	var FoundationProduct models.TProduct

	paramId, err := strconv.Atoi(param_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found", "second_error": err.Error()})
			c.Abort()
			return
		}
	
	if err := db.DB.First(&FoundationProduct, paramId).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found", "second_error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, FoundationProduct)
}

func EditProduct(c *gin.Context){
	var updatedProduct models.TProduct
	
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials json", "second_error": err.Error()})
		return
	}

	if err := db.DB.Save(&updatedProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials sql", "second_error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updatedupdatedProduct", "updated_product": updatedProduct})
}

func DeleteProduct(c *gin.Context){
	id := c.GetInt("param_id")
	
	if err := db.DB.Delete(&models.TProduct{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
