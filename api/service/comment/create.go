package comment

import (
	"fmt"
	"github.com/satori/go.uuid"
	"video/api/model"
)

func (s *commentsService) Create(data *model.Comment) error {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}

	data.ID = u.String()
	return s.dao.Create(data)

}