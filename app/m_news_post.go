package app

import (
	"time"
)

type NewsPost struct {
	Id        int64
	NewsId    int64
	UserId    int64
	Text      string
	IsOline   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (np NewsPost) Save() int64 {
	db.Save(&np)
	if np.Id != 0 {
		return np.Id
	} else {
		db.Last(&np)
		return np.Id
	}
}

// Delete permet de supprimer une catégorie des news
func (np NewsPost) Delete() {
	db.Delete(&np)
}
