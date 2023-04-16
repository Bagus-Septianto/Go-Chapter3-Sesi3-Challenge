package user

import (
	"challenge294/app/models"
	"challenge294/pkg/common"
	"challenge294/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister godoc
//	@Summary		Register User
//	@Description	Creating new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			User	body		models.User	true	"User"
//	@Success		201		{object}	string
//	@Failure		400
//	@Router			/users/register [post]
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