package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"video/api/model"
	"video/api/pkg/errno"
	"video/api/service/comment"
)

// @Description CreateComment
// @Summary CreateComment
// @Tags Comment
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param uid path string true "uid"
// @Param vid path string true "uid"
// @Param content formData string true "content"
// @Resource Create
// @Router /user/{uid}/comment [post]
// @Success 200
func Create(c *gin.Context)  {
	b := &CreateRequest{}

	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	com := &model.Comment{
		UserID: b.UserID,
		VideoID: b.VideoID,
		Content: b.Content,
	}
	fmt.Print(com)

	commentService := comment.NewCommentsService()
	commentService.Create(com)
	c.JSON(200, com)

}
