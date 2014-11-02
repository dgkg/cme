package model

import (
	"log"
	. "strconv"
	"time"
)

type TutorialCategory struct {
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
func (tc TutorialCategory) Save() int64 {
	db.Save(&tc)
	if tc.Id != 0 {
		return tc.Id
	} else {
		db.Last(&tc)
		return tc.Id
	}
}

// Delete permet de supprimer une catégorie d'un tutorial
func (tc TutorialCategory) Delete() {
	db.Delete(&tc)
}

func (tc TutorialCategory) GetTitle() string {
	log.Println("getTitle appelé")
	tc = tc.GetData()

	return tc.Title
}

/*
func (tc *TutorialCategory) getData() {
	if tc.Id != 0 {
		tutoCat := new(TutorialCategory)
		log.Println("getData TutorialCategory appelé")
		idTutoCat := Itoa(int(tc.Id))
		log.Println(idTutoCat)
		db.First(&tutoCat, tc.Id)
		//tc = tutoCat
		//db.Find(&tc)
		log.Println(tutoCat.Title)
	}
	log.Println("Fin : getData TutorialCategory appelé")
}
*/

// permet de récupérer toute la question du forum
// en fonction de son id
func (tc TutorialCategory) GetData() TutorialCategory {
	if tc.Id != 0 {
		db.Where("id = ?", Itoa(int(tc.Id))).Find(&tc)
	}
	return tc
}
