package controllers

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadAllProduct godoc
//	@Summary	Read all Products
//	@Tags		Product
//	@Produce	json
//	@Success	200	{array}	models.Product
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/products [get]
//	@Security	BearerAuth
func (controller ProductController) ReadAllProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	Product := models.Product{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	productReturn, err := controller.ProductService.ReadAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productReturn)
}
