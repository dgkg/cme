package app

import (
	"time"
)

type News struct {
	Id             int64
	UserId         int64
	NewsCategoryId int64
	Title          string `sql:"type:varchar(100);"`
	Text           string
	Keywords       string `sql:"type:varchar(200);"`
	IsOnline       int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Posts          []NewsPost
}
