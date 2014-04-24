package app

import (
	"time"
)

type Tutorial struct {
	Id                 int64
	UserId             int64
	TutorialCategoryId int64
	Title              string `sql:"type:varchar(100);"`
	Text               string
	Keywords           string `sql:"type:varchar(200);"`
	IsOnline           int64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Posts              []TutorialPost
}
