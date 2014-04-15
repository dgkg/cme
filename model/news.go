package model

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

type NewsCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewsPost struct {
	Id        int64
	NewsId    int64
	UserId    int64
	Text      string
	IsOline   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Helper de vue
type NewsViewHelper struct {
	CategoryTitle string
	News
}