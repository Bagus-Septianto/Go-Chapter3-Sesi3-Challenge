package user

import (
	"challenge294/pkg/helpers"
	"challenge294/pkg/common"
	"challenge294/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller UserController) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	
	// err := db.Debug().Where("username = ?", User.Username).Take(&User).Error
	token, err := controller.UserService.Login(User.Username, User.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invaild username/password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}