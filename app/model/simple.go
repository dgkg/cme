package model

import (
	. "strconv"
	"time"
)

// PageSimple est un type qui permet d'afficher des pages d'information
type Simple struct {
	Id        int64
	Url       string
	Title     string
	Text      string
	keywords  string
	IsOnline  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// fonction public
// permet d'enregistrer les éléments du formulaire
func (s Simple) Save() int64 {
	db.Save(&s)
	if s.Id != 0 {
		return s.Id
	} else {
		db.Last(&s)
		return s.Id
	}
}

// fonction public
// permet de supprimer une page simple
func (s Simple) Delete() {
	db.Delete(&s)
}

// permet de récupérer une page simple
// en fonction de son id si il est setup ou de son url
func (s Simple) GetData() Simple {
	if s.Id != 0 {
		db.Where("is_online = ? AND id = ?", "1", Itoa(int(s.Id))).Find(&s)
	} else {
		db.Where("is_online = ? AND url = ?", "1", s.Url).Find(&s)
	}
	return s
}
