package app

import (
	"time"
)

type UserImage struct {
	Id                  int64
	UserId              int64
	UserImageCategoryId int64
	Title               string `sql:"type:varchar(100);"`
	Url                 string `sql:"type:varchar(200);"`
	Description         string
	IsFeatured          int64
	IsOnline            int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
