package app

import (
	"time"
)

type UserProjectCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	Url       string `sql:"_"` // Ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (uic UserProjectCategory) Save() int64 {
	db.Save(&uic)
	if uic.Id != 0 {
		return uic.Id
	} else {
		db.Last(&uic)
		return uic.Id
	}
}

// Delete permet de supprimer une catégorie d'une image
func (uic UserProjectCategory) Delete() {
	db.Delete(&uic)
}

// fonction permettant de récupérer tous les catégories
func (uic UserProjectCategory) getAllCategories() []UserProjectCategory {
	var ups []UserProjectCategory
	db.Find(&ups)
	return ups
}
