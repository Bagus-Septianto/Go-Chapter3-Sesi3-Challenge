package controllers

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
//	@Summary		Create a Product
//	@Description	Creating a Product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Product	body		string			true	"Product Fields"	SchemaExample({"title":"Title of the Product","description":"Lorem ipsum dolor sit amet"})
//	@Success		201		{object}	models.Product	"Product"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/products [post]
//	@Security		BearerAuth
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
