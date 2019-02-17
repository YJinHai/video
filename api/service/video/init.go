package video

import (
	"video/api/dao"
	"video/api/datasource"
	"video/api/model"
)

type VideoService interface {
	Get (id string) (*model.Video)
	Delete (id string) error
	Create(user *model.Video) error
	GetAll(uid string, page, size int) ([]*model.Video, error)
	Comments(vid string, page, limit int) []*model.Comment
}

type videoService struct {
	dao *dao.VideoDao
}

func NewVideoService() VideoService {
	return &videoService{
		dao: dao.NewVideoDao(datasource.InstanceDbMaster()),
	}
}
