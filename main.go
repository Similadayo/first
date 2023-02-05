package main

import (
	"fmt"

	"github.com/Similadayo/db"
	"github.com/Similadayo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDB()
	defer db.DB.Close()

	//route setup
	r := gin.Default()
	routes.InitializeRoutes(r)
	fmt.Println("server is running...")
	r.Run(":8080")
}
