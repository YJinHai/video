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
	VideoID string `gorm:"column:video_id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (v *VideoDelrec) TableName() string {
	return "video_delrec"
}
