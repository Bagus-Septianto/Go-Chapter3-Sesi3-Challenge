package controllers

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller ProductController) ReadProduct(c *gin.Context) {
	// Kalau sudah masuk sini. product.UserID yang dicari udah sama dengan UserID yang terautentikasi
	// query 2x sih jadinya... (1x di middleware.authorization, 1x lagi di repository)
	// males ngemodif middleware hehe (if no brokie don't touch it) -bagus
	contentType := helpers.GetContentType(c)

	Product := models.Product{} // makin malem makin sadar, ngapain ngebind body/json product kalo ngeread...

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	productReturn, err := controller.ProductService.Read(productId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productReturn)
}
