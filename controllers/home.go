package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to First-Clost")
}
