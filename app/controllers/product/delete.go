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

// DeleteProduct godoc
//	@Summary		Delete a Product
//	@Description	Deleting a Product by ProductID
//	@Tags			Product
//	@Produce		json
//	@Param			productId	path	int	true	"Product ID"
//	@Success		204
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/products/{productId} [delete]
//	@Security		BearerAuth
func (controller ProductController) DeleteProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := controller.ProductService.Delete(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
