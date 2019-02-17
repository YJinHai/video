package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"video/api/model"
)

type CommentDao struct {
	engine *gorm.DB
}

func NewCommentDao(engine *gorm.DB) *CommentDao {
	return &CommentDao{
		engine: engine,
	}
}

func (d *CommentDao) Create(data *model.Comment) error {

	return d.engine.Create(data).Error
}

func (d *CommentDao) Get (id string) ([]*model.Comment, error) {
	//vc := &model.Comment{}
	c:= make([]*model.Comment, 0)
	//u := model.User{ID:id}
	//u.ID = id
	//if err := d.engine.Where("id=?", id).Preload("User").Preload("Video").First(c).Error; err != nil {
	//	return c, err
	//}
	d.engine.Where("user_id=?", id).Find(&c)

	//d.engine.Where("id=?", id).Find(u)
	//fmt.Print(u)


	//d.engine.Model(u).Association("User").Find(&c)
	//d.engine.Model(&u).Related(&c).Find(&c)
	//d.engine.First(u, id).Related(&c).Find(&c)

	//d.engine.Where("user_id = ?", id).Find(c)
	//d.engine.Where("id = ?", id).First(c)
	fmt.Print(c)
	return c, nil
}

func (d *CommentDao) Delete (id string) error {
	u := &model.Comment{ID:id}
	return d.engine.Delete(u).Error
}

func (d *CommentDao) GetAll(uid int, page, limit int) ([]*model.Comment, error) {
	offset := (page - 1) * limit
	comments := make([]*model.Comment, 0)

	if err := d.engine.Where("user_id = ?", uid).Offset(offset).Limit(limit).Order("add_time desc").Find(comments).Error; err != nil {
		return comments, err
	}

	return comments, nil

}
