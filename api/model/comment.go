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

type Comment struct {
	ID      string         `gorm:"column:id;primary_key"`
	VideoID string         `gorm:"column:video_id"`
	UserID  string         `gorm:"column:user_id"`
	Content string 		   `gorm:"column:content"`
	AddTime time.Time      `gorm:"column:add_time"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comments"
}
