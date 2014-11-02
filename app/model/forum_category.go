package model

import (
	"time"
)

type ForumCategory struct {
	Id         int64
	Title      string `sql:"type:varchar(100);"`
	Url        string `sql:"_"` // Ignore this field
	IsSelected bool   `sql:"_"` // Ignore this field
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (fc ForumCategory) Save() int64 {
	db.Save(&fc)
	if fc.Id != 0 {
		return fc.Id
	} else {
		db.Last(&fc)
		return fc.Id
	}
}

// Delete permet de supprimer une catégorie du forum
func (fc ForumCategory) Delete() {
	db.Delete(&fc)
}
