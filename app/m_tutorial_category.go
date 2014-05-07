package app

import (
	"log"
	. "strconv"
	"time"
)

type TutorialCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	Url       string `sql:"_"` // Ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (tc TutorialCategory) getTitle() string {
	log.Println("getTitle appelé")
	tc = tc.getData()

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
func (tc TutorialCategory) getData() TutorialCategory {
	if tc.Id != 0 {
		db.Where("id = ?", Itoa(int(tc.Id))).Find(&tc)
	}
	return tc
}
