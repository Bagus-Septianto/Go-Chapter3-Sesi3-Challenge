package main

import (
	router "challenge294/app/controllers"
	userRepo "challenge294/app/repository/user"
	userCont "challenge294/app/controllers/user"
	userServ "challenge294/app/services/user"
	productRepo "challenge294/app/repository/product"
	productCont "challenge294/app/controllers/product"
	productServ "challenge294/app/services/product"
	"challenge294/pkg/database"
)

// https://softwareengineering.stackexchange.com/questions/306890/is-it-bad-practice-that-a-controller-calls-a-repository-instead-of-a-service

func main() {
	db, err := database.StartDB()
	if err != nil {
		panic(err)
	}
	userRepository := userRepo.NewUserRepository()
	userService := userServ.NewUserService(userRepository, db)
	userController := userCont.NewUserController(userService, db)
	
	productRepository := productRepo.NewProductRepository()
	productService := productServ.NewProductService(productRepository, db)
	productController := productCont.NewProductController(productService, db)

	r := router.StartApp(userController, productController)
	r.Run(":8080")
}