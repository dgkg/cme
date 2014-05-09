package app

import (
	"time"
)

type ForumCategory struct {
	Id         int64
	Title      string `sql:"type:varchar(100);"`
	Url        string `sql:"_"` // Ignore this field
	IsSelected bool   `sql:"_"` // Ignore this field
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
