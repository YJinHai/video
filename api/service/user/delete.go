package user

func (s *userService) Delete (id string) error {
	return s.dao.Delete(id)
}
