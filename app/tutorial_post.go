package app

import (
	"time"
)

type TutorialPost struct {
	Id         int64
	TutorialId int64
	UserId     int64
	Text       string
	IsOline    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
