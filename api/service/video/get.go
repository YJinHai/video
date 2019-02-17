package video

import "video/api/model"

func (s *videoService) Get (id string) (*model.Video) {
	//data := s.getByCache(id)
	//if data == nil || data.ID <= 0 {
	//	data = s.dao.Get(id)
	//	if data == nil || data.ID <= 0 {
	//		data = &model.Video{ID: id}
	//	}
	//	s.setByCache(data)
	//}
	data, _ := s.dao.Get(id)
	return data
}
