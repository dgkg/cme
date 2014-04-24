package app

import (
	"time"
)

type ForumPost struct {
	Id        int64
	ForumId   int64
	UserId    int64
	Text      string
	IsOline   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
