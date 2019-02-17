package video

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"video/api/service/video"
)

// @Description VideoCommentsList
// @Summary VideoCommentsList
// @Tags Video
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param uid path int true "userID"
// @Param vid path int true "videoID"
// @Param page formData int true "page"
// @Param limit formData int true "limit"
// @Resource Comments
// @Router /user/{uid}/video/{vid}/comments [post]
// @Success 200 {object} model.Comment
func Comments(c *gin.Context)  {
	id := c.Param("vid")
	page, _ := strconv.Atoi(c.PostForm("page"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	videoService := video.NewVideoService()
	c.JSON(200, videoService.Comments(id, page, limit))
}
