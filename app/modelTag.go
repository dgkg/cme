package controler

import (
	"time"
)

type Tag struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	IsOnline  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
