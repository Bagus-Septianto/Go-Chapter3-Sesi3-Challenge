package controllers

import (
	"challenge294/app/services"

	"gorm.io/gorm"
)

type ProductController struct {
	ProductService services.IProductService
	DB          *gorm.DB
}


func NewProductController(productService services.IProductService, db *gorm.DB) *ProductController {
	return &ProductController{
		ProductService: productService,
		DB: db,
	}
}
