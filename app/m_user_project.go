package app

import (
	//"fmt"
	. "strconv"
	//"strings"
	"time"
)

type UserProject struct {
	Id                    int64
	UserId                int64
	UserProjectCategoryId int64
	ImageUrl              string
	Title                 string `sql:"type:varchar(100);"`
	Description           string
	IsFeatured            int64
	IsOnline              int64
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Images                []UserImage
}

// fonction public
// permet d'enregistrer les éléments du projet d'un user
// retourne l'id du projet
func (up UserProject) Save() int64 {
	db.Save(&up)
	if up.Id != 0 {
		return up.Id
	} else {
		db.Last(&up)
		return up.Id
	}
}

// fonction public
// permet de supprimer un projet d'un user
func (up UserProject) Delete() {
	db.Delete(&up)
}

// getById permet de récupérer toute les données d'un projet de user
// en fonction de son id
func (up UserProject) getById() UserProject {
	db.Where("is_online = ? AND id = ?", "1", Itoa(int(up.Id))).Find(&up)
	return up
}

// getAllFromIdUser permet de récupérer tous les projets d'un user
// en fonction de l'id user
func (up UserProject) getAllFromIdUser() []UserProject {
	var ups []UserProject
	db.Where("is_online = ? AND user_id = ?", "1", Itoa(int(up.UserId))).Find(&ups)
	return ups
}
