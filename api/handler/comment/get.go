package comment

import (
	"github.com/gin-gonic/gin"
	"video/api/service/comment"
)

// @Description GetComment
// @Summary GetComment
// @Accept json
// @Tags Comment
// @Security Bearer
// @Produce  json
// @Param id path int true "ID"
// @Resource Get
// @Router /comment/{id} [get]
// @Success 200 {object} model.Comment
func Get(c *gin.Context) {
	id := c.Param("id")
	commentsService := comment.NewCommentsService()
	c.JSON(200, commentsService.Get(id))
}
