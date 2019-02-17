package video

import "video/api/model"

func (s *videoService) Comments(vid string, page, limit int) []*model.Comment {
	data,_ := s.dao.VideoCommentsList(vid, page, limit)
	return data
}
