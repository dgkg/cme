package model

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

type TutorialCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TutorialPost struct {
	Id         int64
	TutorialId int64
	UserId     int64
	Text       string
	IsOline    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TutorialsViewHelper struct {
	Tutorial
}