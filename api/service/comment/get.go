package comment

import "video/api/model"

func (s *commentsService) Get (id string) []*model.Comment {
	//data := s.getByCache(id)
	//if data == nil || data.ID <= 0 {
	//	data, _ = s.dao.Get(id)
	//	if data == nil || data.ID <= 0 {
	//		data = &model.Comment{Id: id}
	//	}
	//	s.setByCache(data)
	//}
	data, _ := s.dao.Get(id)
	return data
}
