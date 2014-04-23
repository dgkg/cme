package app

import (
	"time"
)

type User struct {
	Id        int64
	FirstName string `sql:"type:varchar(100);"`
	LastName  string `sql:"type:varchar(100);"`
	Text      string
	Email     string
	Pass      string
	Keywords  string
	IsOnline  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Images    []UserImage
}
