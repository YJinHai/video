package video

import "video/api/model"

func (s *videoService) GetAll(uid string, page, size int) ([]*model.Video, error) {
	return s.dao.GetAll(uid, page, size)
}
