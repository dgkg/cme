package app

import (
	"time"
)

type NewsCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
