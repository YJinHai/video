package user

import (
	"fmt"
	"github.com/satori/go.uuid"
	"video/api/model"
)

func (s *userService) Create(data *model.User) error {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}

	data.ID = u.String()
	return s.dao.Create(data)
}
