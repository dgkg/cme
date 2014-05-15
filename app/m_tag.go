package app

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

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (t Tag) Save() int64 {
	db.Save(&t)
	if t.Id != 0 {
		return t.Id
	} else {
		db.Last(&t)
		return t.Id
	}
}

// Delete permet de supprimer un tag
func (t Tag) Delete() {
	db.Delete(&t)
}
