package controllers

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (controller ProductController) CreateProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}
	
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt, ini di controller apa service gatau
	userID := uint(userData["id"].(float64))

	productReturn, err := controller.ProductService.Create(&Product, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, productReturn)
}
