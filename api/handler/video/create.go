package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"video/api/model"
	"video/api/pkg/errno"
	"video/api/service/video"
)

// @Description CreateVideo
// @Summary CreateVideo
// @Tags Video
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param uid path string true "uid"
// @Param title formData string true "name"
// @Param dc formData string true "display_ctime"
// @Resource Create
// @Router /user/{uid}/video [post]
// @Success 200
func Create(c *gin.Context)  {

	//fmt.Print(c.PostForm("dc"))
	b := &CreateRequest{}
	if err := c.ShouldBind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	v := &model.Video{
		UserID: b.UserID,
		Title: b.Title,
		DisplayCtime: c.PostForm("dc"),
	}

	fmt.Print("video:", v)

	videoService := video.NewVideoService()
	videoService.Create(v)
	c.JSON(200, v)

}
