package video

func (s *videoService) Delete (id string) error {
	return s.dao.Delete(id)
}
