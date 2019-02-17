package user

import "video/api/model"

func (s *userService) Videos(uid string, page, limit int) []*model.Video {
	data,_ := s.dao.UserVideosList(uid, page, limit)
	return data
}
