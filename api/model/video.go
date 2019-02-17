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

type Video struct {
	ID           string         `gorm:"column:id;primary_key"`
	Title        string         `gorm:"column:title"`
	UserID       string         `gorm:"column:user_id"`
	DisplayCtime string 		`gorm:"column:display_ctime"`
	AddTime      time.Time      `gorm:"column:add_time"`
}

// TableName sets the insert table name for this struct type
func (v *Video) TableName() string {
	return "videos"
}
