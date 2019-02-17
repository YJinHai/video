package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"video/api/service/user"
)

// @Description GetVideosList
// @Summary GetVideosList
// @Tags User
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param uid path int true "nameID"
// @Param page formData int true "page"
// @Param limit formData int true "limit"
// @Resource Videos
// @Router /user/{uid}/videos [post]
// @Success 200 {object} model.Video
func Videos(c *gin.Context) {
	id := c.Param("uid")
	page, _ := strconv.Atoi(c.PostForm("page"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	userService := user.NewUserService()
	c.JSON(200, userService.Videos(id, page, limit))
}
