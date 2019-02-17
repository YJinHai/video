package user

import (
	"video/api/dao"
	"video/api/datasource"
	"video/api/model"
)

type UserService interface {
	Get (id string) (*model.User)
	Delete (id string) error
	Create(user *model.User) error
	Comments(uid string, page, limit int) []*model.Comment
	Videos(uid string, page, limit int) []*model.Video
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}
