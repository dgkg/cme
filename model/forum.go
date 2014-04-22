package model

import (
	"time"
)

type Forum struct {
	Id              int64
	UserId          int64
	ForumCategoryId int64
	Title           string `sql:"type:varchar(100);"`
	Text            string
	Keywords        string `sql:"type:varchar(200);"`
	IsSolved        int64
	IsOnline        int64
	PostNumb        int64  `sql:"-"` // Ignore this field
	CategoryTitle   string `sql:"-"` // Ignore this field
	CreatedAt       time.Time
	CreatedAtString string `sql:"-"` // Ignore this field
	UpdatedAt       time.Time
	Posts           []ForumPost
	Users           []User `sql:"-"` // Ignore this field
}

type ForumCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	Url       string `sql:"_"` // Ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ForumPost struct {
	Id        int64
	ForumId   int64
	UserId    int64
	Text      string
	IsOline   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Helper de vue
type ForumViewHelper struct {
	CategoryTitle string
	Forum
}
