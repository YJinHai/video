package video

import (
	"github.com/gin-gonic/gin"
	"video/api/service/video"
)

// GetVideo
// @Summary GetVideo
// @Accept json
// @Tags Video
// @Security Bearer
// @Produce  json
// @Param uname path string true "nameID"
// @Resource Get
// @Router /user/{uid}/videos/{vid-id} [get]
// @Success 200 {object} model.Video
func Get(c *gin.Context) {
	vid := c.Param("vid-id")
	videoService := video.NewVideoService()
	c.JSON(200, videoService.Get(vid))

}
