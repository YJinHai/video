package user

import "video/api/model"

func (s *userService) Get (id string) (*model.User) {
	//data := s.getByCache(id)
	//if data == nil || data.ID <= 0 {
	//	data = s.dao.Get(id)
	//	if data == nil || data.Id == "0" {
	//		data = &model.User{Id: id}
	//	}
	//	s.setByCache(data)
	//}
	data,_ := s.dao.Get(id)
	return data
}

