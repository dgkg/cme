package app

import (
	"time"
)

type UserImageCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
