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

type User struct {
	ID       string `gorm:"column:id;primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Comments []Comment  `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
	Videos []Video  `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
