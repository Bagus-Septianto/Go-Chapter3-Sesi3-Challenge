package controllers

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UpdateProduct godoc
//	@Summary		Update a Product
//	@Description	Updating a Product by ProductID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			productId	path		int				true	"Product ID"
//	@Param			Product		body		string			true	"Product Fields"	SchemaExample({"title":"Title of the Product","description":"Lorem ipsum dolor sit amet"})
//	@Success		200			{object}	models.Product	"Product"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/products/{productId} [put]
//	@Security		BearerAuth
func (controller ProductController) UpdateProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt
	userID := uint(userData["id"].(float64))

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	productReturn, err := controller.ProductService.Update(&Product, productId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productReturn)
}