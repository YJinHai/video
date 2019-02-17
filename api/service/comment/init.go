package comment

import (
	"video/api/dao"
	"video/api/datasource"
	"video/api/model"
)

type CommentsService interface {
	Get (id string) []*model.Comment
	Delete (id string) error
	Create(comments *model.Comment) error
	//GetAll(page, size int) []model.Comment
}

type commentsService struct {
	dao *dao.CommentDao
}


func NewCommentsService() CommentsService {
	return &commentsService{
		dao: dao.NewCommentDao(datasource.InstanceDbMaster()),
	}
}
