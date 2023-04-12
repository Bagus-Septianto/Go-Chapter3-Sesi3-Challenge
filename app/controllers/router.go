package controllers

import (
	// "challenge294/app/controllers/user"
	// product "challenge294/app/controllers/product"

	"challenge294/app/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(userControllers IUserControllers, productController IProductControllers) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userControllers.UserRegister)

		userRouter.POST("/login", userControllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication()) // Punya token yang valid ? lanjut : keluar

		productRouter.POST("/", productController.CreateProduct) // Create, produk yang dibuat akan langsung diset dengan id user yang telah terautentikasi
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), productController.ReadProduct) // Read Product By ID
		productRouter.PUT("/:productId", middlewares.AdminOnlyAuthorization(), productController.UpdateProduct) // Update
		productRouter.DELETE("/:productId", middlewares.AdminOnlyAuthorization(), productController.DeleteProduct) // Delete
		productRouter.GET("/", middlewares.AdminOnlyAuthorization(), productController.ReadAllProduct) // Read Product By ID
	}

	return r
}