package app

import (
	"time"
)

type ForumPost struct {
	Id              int64
	ForumId         int64
	UserId          int64
	Text            string
	IsOnline         int64
	UserName        string `sql:"-"` // Ignore this field
	CreatedAtString string `sql:"-"` // Ignore this field
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// fonction public
// permet d'enregistrer les éléments du formulaire
func (fp ForumPost) Save() {
	db.Save(&fp)
}
