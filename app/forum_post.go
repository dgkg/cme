package app

import (
	"time"
)

type ForumPost struct {
	Id              int64
	ForumId         int64
	UserId          int64
	Text            string
	IsOline         int64
	UserName        string `sql:"-"` // Ignore this field
	CreatedAtString string `sql:"-"` // Ignore this field
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
