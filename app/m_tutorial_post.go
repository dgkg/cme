package app

import (
	. "strconv"
	"time"
)

type TutorialPost struct {
	Id         int64
	TutorialId int64
	UserId     int64
	Text       string
	IsOline    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (tp TutorialPost) Save() int64 {
	db.Save(&tp)
	if tp.Id != 0 {
		return tp.Id
	} else {
		db.Last(&tp)
		return tp.Id
	}
}

// Delete permet de supprimer un tutorial post
func (tp TutorialPost) Delete() {
	db.Delete(&tp)
}

func (t *TutorialPost) getData() {
	if t.Id != 0 {
		db.Where("id=?", Itoa(int(t.Id))).Find(&t)
	}
}
