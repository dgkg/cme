package app

import (
	"time"
)

type User struct {
	Id        int64
	FirstName string `sql:"type:varchar(100);"`
	LastName  string `sql:"type:varchar(100);"`
	Text      string
	Email     string
	Pass      string
	Keywords  string
	IsOnline  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Images    []UserImage
}

func (u User) getById() User {
	var user User
	user.Id = u.Id
	db.Find(&user)
	return user
}

/*
func (u *User) getById() User {
	return db.Find(&u)
}
*/
