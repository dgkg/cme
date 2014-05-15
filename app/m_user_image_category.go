package app

import (
	"time"
)

type UserImageCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	Url       string `sql:"_"` // Ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (uic UserImageCategory) Save() int64 {
	db.Save(&uic)
	if uic.Id != 0 {
		return uic.Id
	} else {
		db.Last(&uic)
		return uic.Id
	}
}

// Delete permet de supprimer une catégorie d'une image
func (uic UserImageCategory) Delete() {
	db.Delete(&uic)
}
