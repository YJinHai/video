package user

import "video/api/model"

func (s *userService) Comments(uid string, page, limit int) []*model.Comment {
	data,_ := s.dao.UserCommentsList(uid, page, limit)
	return data
}
