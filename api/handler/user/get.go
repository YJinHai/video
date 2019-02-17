package user

import (
	"github.com/gin-gonic/gin"
	"video/api/service/user"
)

// @Description GetUser
// @Summary Get
// @Accept json
// @Tags User
// @Security Bearer
// @Produce  json
// @Param uid path string true "nameID"
// @Resource Get
// @Router /user/{uid} [get]
// @Success 200 {object} model.User
func Get(c *gin.Context) {
	uid := c.Param("uid")
	userService := user.NewUserService()
	c.JSON(200, userService.Get(uid))
}
