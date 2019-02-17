package comment

func (s *commentsService) Delete (id string) error {
	return s.dao.Delete(id)
}

