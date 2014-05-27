package model

import (
	"time"
)

type NewsCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	Url       string `sql:"_"` // Ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (nc NewsCategory) Save() int64 {
	db.Save(&nc)
	if nc.Id != 0 {
		return nc.Id
	} else {
		db.Last(&nc)
		return nc.Id
	}
}

// Delete permet de supprimer une catégorie des news
func (nc NewsCategory) Delete() {
	db.Delete(&nc)
}
