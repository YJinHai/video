package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"video/api/model"
)

type VideoDao struct {
	engine *gorm.DB
}

func NewVideoDao(engine *gorm.DB) *VideoDao {
	return &VideoDao{
		engine: engine,
	}
}

func (d *VideoDao) Create(data *model.Video) error {
	u := &model.User{}
	d.engine.Where("username=?", data.UserID).First(u)
	fmt.Print("11data.userid:",data.UserID)
	data.UserID = u.ID
	fmt.Print("u.id:",u.ID)
	fmt.Print("data.userid:",data.UserID)
	return d.engine.Create(data).Error
}

func (d *VideoDao) Get (id string) (*model.Video, error) {
	u := &model.Video{}
	data :=d.engine.Where("id = ?", id).First(u)
	return u, data.Error
}

func (d *VideoDao) Delete (id string) error {
	u := &model.Video{ID:id}
	return d.engine.Delete(u).Error
}

func (d *VideoDao) GetAll(uid string, page, limit int) ([]*model.Video, error) {
	offset := (page - 1) * limit
	comments := make([]*model.Video, 0)

	if err := d.engine.Where("user_id = ?", uid).Offset(offset).Limit(limit).Order("add_time desc").Find(comments).Error; err != nil {
		return comments, err
	}

	return comments, nil

}

func (d *VideoDao) VideoCommentsList(vid string, page, limit int) ([]*model.Comment, error) {
	offset := (page - 1) * limit
	comments := make([]*model.Comment, 0)

	if err := d.engine.Where("video_id = ?", vid).Offset(offset).Limit(limit).Order("add_time desc").Find(&comments).Error; err != nil {
		return comments, err
	}

	return comments, nil

}
