package app

import (
	. "strconv"
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

func (t *TutorialPost) getData() {
	if t.Id != 0 {
		db.Where("id=?", Itoa(int(t.Id))).Find(&t)
	}
}
