package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"video/api/model"
)

type UserDao struct {
	engine *gorm.DB
}

func NewUserDao(engine *gorm.DB) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (d *UserDao) Create(data *model.User) error {
	return d.engine.Create(data).Error
}

func (d *UserDao) Get (id string) (*model.User, error) {
	u := &model.User{}
	data :=d.engine.Where("id = ?", id).First(u)
	return u, data.Error
}

func (d *UserDao) Delete (id string) error {
	u := &model.User{ID:id}
	return d.engine.Delete(u).Error
}

func (d *UserDao) UserCommentsList(uid string, page, limit int) ([]*model.Comment, error) {
	offset := (page - 1) * limit
	comments := make([]*model.Comment, 0)
	user := &model.User{
		ID: uid,
	}

	//if err := d.engine.Where("user_id = ?", uid).Offset(offset).Limit(limit).Order("add_time desc").Find(&comments).Error; err != nil {
	//	fmt.Print("errrror")
	//	return comments, err
	//}
	if err := d.engine.Model(user).Related(&comments).Offset(offset).Limit(limit).Order("add_time desc").Find(&comments).Error; err != nil {
		fmt.Print("errrror")
		return comments, err
	}

	return comments, nil

}

func (d *UserDao) UserVideosList(uid string, page, limit int) ([]*model.Video, error) {
	offset := (page - 1) * limit
	videos := make([]*model.Video, 0)
	user := &model.User{
		ID: uid,
	}

	if err := d.engine.Model(user).Related(&videos).Offset(offset).Limit(limit).Order("add_time desc").Find(&videos).Error; err != nil {
		fmt.Print("errrror")
		return videos, err
	}

	return videos, nil

}
