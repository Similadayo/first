package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Similadayo/db"
	"github.com/Similadayo/models"
	"github.com/Similadayo/utils"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var product models.Product
	var err error
	if err = db.DB.Preload("Category").First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	db.DB.Preload("Category").Find(&products)
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	// Get the token from the Authorization header
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in request headers"})
		return
	}

	// Verify the token
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the user is an admin
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can access this resource"})
		return
	}

	// Get the product information from the request body
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	// Link the category with the categoryID
	var category models.Category
	if err := db.DB.Where("id = ?", product.CategoryID).First(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding category in database"})
		log.Println(err)
		return
	}

	product.Category = category

	// Create the product in the database
	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	// Return the created product
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func GetProductsByCategory(c *gin.Context) {
	// Get the category ID from the request parameters
	categoryID := c.Param("id")

	// Get the products by category
	var products []models.Product
	if err := db.DB.Preload("Category").Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
		return
	}

	// Return the products
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func UpdateProduct(c *gin.Context) {
	// Get the token from the Authorization header
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in request headers"})
		return
	}

	// Verify the token
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the user is an admin
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can access this resource"})
		return
	}

	// Get the product ID from the request parameters
	productID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	// Get the product from the database
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to database"})
		return
	}
	defer db.Close()

	var product models.Product
	if err := db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	// Bind the updated product data to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	// Save the updated product to the database
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Category updated successfully",
	})
}

func DeleteProduct(c *gin.Context) {
	// Get the token from the Authorization header
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in request headers"})
		return
	}

	// Verify the token
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the user is an admin
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can access this resource"})
		return
	}

	// Get the category ID from the request parameters
	productID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Get the category from the database
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message(false, "Error connecting to database"))
		return
	}
	defer db.Close()

	var product models.Product
	if err := db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Delete the category from the database
	if err := db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message(false, "Error deleting category"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Category deleted successfully",
	})
}
