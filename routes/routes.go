package routes

import (
	"github.com/Similadayo/controllers"
	"github.com/Similadayo/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomeController)

	//Register and login
	userController := r.Group("/")
	{
		userController.POST("register", controllers.CreateUser)
		userController.POST("login", controllers.Login)
		userController.POST("logout", controllers.Logout)

		userController.GET("getusers", middleware.Authorization, controllers.GetUsers)
		userController.GET("getuser/:id", middleware.Authorization, controllers.GetUser)

		userController.POST("updatepassword/:id", middleware.Authorization, controllers.ChangePassword)
		userController.POST("forgotpassword", controllers.ForgotPassword)

		userController.PUT("updateuser/:id", middleware.Authorization, controllers.UpdateUser)
		userController.DELETE("deactivate/:id", middleware.Authorization, controllers.DeleteUser)
		userController.PUT("suspend/:id", middleware.Authorization, controllers.SuspendUser)
	}

	//Category routes
	categoryController := r.Group("/category")
	{
		categoryController.GET("/", middleware.Authorization, controllers.GetCategories)
		categoryController.GET("/:id", middleware.Authorization, controllers.GetCategory)
		categoryController.POST("/create", middleware.Authorization, controllers.CreateCategory)
		categoryController.PUT("/update/:id", middleware.Authorization, controllers.UpdateCategory)
		categoryController.DELETE("/delete/:id", middleware.Authorization, controllers.DeleteCategory)
	}

	// Product routes
	productController := r.Group("/product")
	{
		productController.GET("/", middleware.Authorization, controllers.GetProducts)
		productController.GET("/:id", middleware.Authorization, controllers.GetProduct)
		productController.GET("/category/:id", middleware.Authorization, controllers.GetProductsByCategory)
		productController.POST("/createproduct", middleware.Authorization, controllers.CreateProduct)
		productController.PUT("/updateproduct/:id", middleware.Authorization, controllers.UpdateProduct)
		productController.DELETE("/deleteproduct/:id", middleware.Authorization, controllers.DeleteProduct)
		productController.GET("/search", middleware.Authorization, controllers.SearchProducts)
	}

	orderController := r.Group("/order")
	{
		orderController.GET("/:id", middleware.Authorization, controllers.GetOrderByID)
		orderController.GET("/", middleware.Authorization, controllers.GetOrders)
		orderController.POST("/new", middleware.Authorization, controllers.CreateOrder)
		orderController.PUT("/update/:id", middleware.Authorization, controllers.UpdateOrder)
		orderController.DELETE("/delete/:id", middleware.Authorization, controllers.DeleteOrder)

	}
}
