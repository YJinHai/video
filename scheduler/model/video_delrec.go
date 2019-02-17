package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type VideoDelrec struct {
	VideoID int `gorm:"column:video_id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (v *VideoDelrec) TableName() string {
	return "video_delrec"
}

func (d *VideoDelrec) Delete(id int) error{
	v := &VideoDelrec{VideoID:id}
	return DB.Self.Delete(v).Error
}

func (d *VideoDelrec) Create(data *VideoDelrec) error {
	return DB.Self.Create(data).Error
}

func (d *VideoDelrec) Get(id int) (*VideoDelrec, error) {
	v := &VideoDelrec{}
	data := DB.Self.Where("VideoID = ?", id).First(v)
	return v, data.Error
}