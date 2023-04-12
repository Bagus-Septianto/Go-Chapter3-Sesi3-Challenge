package user

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func UserRegister(c *gin.Context) {
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(c)
// 	_,_ = db, contentType
// 	User := models.User{}

// 	if contentType == common.AppJSON {
// 		c.ShouldBindJSON(&User)
// 	} else {
// 		c.ShouldBind(&User)
// 	}

// 	err := db.Debug().Create(&User).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Bad Request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"id": User.ID,
// 		"username": User.Username,
// 		"role": User.Role,
// 	})
// }

func (controller UserController) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	// apakah user harus dibalikin dari userservice?
	// mungkin kedepan bisa ada perubahan/balikan id (Ohhhh)
	userReturn, err := controller.UserService.Register(&User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": userReturn.ID,
		"username": userReturn.Username,
		"role": userReturn.Role,
	})
}