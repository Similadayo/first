package controllers

import (
	"log"
	"net/http"

	"github.com/Similadayo/db"
	"github.com/Similadayo/models"
	"github.com/Similadayo/utils"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Println(err)
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the token from the Authorization header
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to make this request"})
		return
	}

	// Verify the token
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	order.UserID = int(claims.UserID)
	user := &models.User{}
	if err := db.First(user, order.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	order.Address = user.Address
	order.Status = "pending"

	order.TotalCost = 0
	for _, product := range order.Products {
		p := &models.Product{}
		if err := db.First(p, product.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}
		order.TotalCost += p.Price
	}

	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

func GetOrderByID(c *gin.Context) {
	db, err := db.GetDB()
	{
		if err != nil {
			log.Println(err)
			return
		}
	}
	orderID := c.Param("id")
	order := &models.Order{}

	if err := db.Preload("Products").First(order, orderID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrders(c *gin.Context) {
	db, err := db.GetDB()
	{
		if err != nil {
			log.Println(err)
			return
		}
	}
	orders := []models.Order{}

	if err := db.Preload("Products").Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Orders not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func UpdateOrder(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Println(err)
		return
	}

	orderID := c.Param("id")
	var order models.Order
	if err := db.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	// Get the token from the Authorization header
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to make this request"})
		return
	}

	// Verify the token
	claims, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the user is an administrator
	if claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to make this request"})
		return
	}

	var updateOrder models.Order
	if err := c.ShouldBindJSON(&updateOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateOrder.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status cannot be empty"})
		return
	}

	// Update the order status
	order.Status = updateOrder.Status
	if err := db.Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": order})
}

func DeleteOrder(c *gin.Context) {
	db, err := db.GetDB()
	{
		if err != nil {
			log.Println(err)
			return
		}
	}
	order := &models.Order{}

	if err := db.First(order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := db.Delete(order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
