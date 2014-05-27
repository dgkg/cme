package model

import (
	"time"
)

type ForumPost struct {
	Id              int64
	ForumId         int64
	UserId          int64
	Text            string
	IsOnline        int64
	UserName        string `sql:"-"` // Ignore this field
	CreatedAtString string `sql:"-"` // Ignore this field
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// fonction public
// permet d'enregistrer les éléments du formulaire
func (fp ForumPost) Save() int64 {
	db.Save(&fp)
	if fp.Id != 0 {
		return fp.Id
	} else {
		db.Last(&fp)
		return fp.Id
	}
}

// fonction public
// permet de supprimer un post sur une question du forum
func (fp ForumPost) Delete() {
	db.Delete(&fp)
}
