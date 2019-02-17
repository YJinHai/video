package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
	"video/api/model"
	"video/api/pkg/errno"
	"video/api/service/user"
)


// @Description CreateUser
// @Summary CreateUser
// @Tags User
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param username formData string true "name"
// @Param pwd formData string true "user pwd"
// @Resource Create
// @Router /user [post]
// @Success 200
func Create(c *gin.Context)  {
	b := &CreateRequest{}

	var err error
	if err := c.Bind(b);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.ErrBind})
		return
	}

	u := &model.User{
		Username: b.Username,
		Password: b.Password,
	}

	log.Debugf("username is: [%s], password is [%s]", u.Username, u.Password)

	if u.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if u.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)

	userService := user.NewUserService()
	userService.Create(u)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

}