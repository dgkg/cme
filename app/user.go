package app

import (
	"errors"
	"log"
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
func (u *User) LoginPassExist() (v bool, err error) {

	log.Println("SearchUserLoginAndPass appelé")
	log.Println(u.Email + " / " + u.Pass)

	var user User
	if u.Email != "" && u.Pass != "" {
		db.Where("email = ? AND pass = ?", u.Email, u.Pass).Find(&user)
	} else {
		err = errors.New("Il manque le login et/ou le mot de passe")
	}

	u.Id = user.Id
	u.LastName = user.LastName
	u.FirstName = user.FirstName

	if u.Id != 0 && u.LastName != "" && u.FirstName != "" {
		v = true
	}

	return
}

/*
func (u *User) getById() User {
	return db.Find(&u)
}
*/
