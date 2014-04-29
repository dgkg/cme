package app

import (
	"errors"
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

// fonction permettant de retourner un utilisateur
// en fonction de son login et son mot de passe
// retourne un utilisateur si trouvé
// renvoie une erreur le cas échéant
func (u User) serchUserLoginAndPass() (user User, err error) {

	if u.Email != "" && u.Pass != "" {
		db.Where("email = ? AND pass = ?", u.Email, u.Pass).Find(&user)
	} else {
		err = errors.New("Il manque le login et/ou le mot de passe")
	}
	return
}

/*
func (u *User) getById() User {
	return db.Find(&u)
}
*/
