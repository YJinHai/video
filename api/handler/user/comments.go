package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"video/api/service/user"
)

// @Description GetCommentsList
// @Summary GetCommentsList
// @Tags User
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param uid path int true "nameID"
// @Param page formData int true "page"
// @Param limit formData int true "limit"
// @Resource Comments
// @Router /user/{uid}/comments [post]
// @Success 200 {object} model.Comment
func Comments(c *gin.Context) {
	//uname := c.Param("uname")
	//page := c.Param("page")
	//limit := c.Param("limit")
	id := c.PostForm("uid")
	page, _ := strconv.Atoi(c.PostForm("page"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	fmt.Print("ss:",id)
	userService := user.NewUserService()
	c.JSON(200, userService.Comments(id, page, limit))
}
