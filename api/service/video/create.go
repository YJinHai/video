package video

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
	"video/api/model"
)

func (s *videoService) Create(data *model.Video) error {
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}

	data.ID = u.String()
	data.AddTime= time.Now()
	return s.dao.Create(data)
}
