package app

import (
	//"fmt"
	. "strconv"
	//"strings"
	"math/rand"
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
	Url                   string `sql:"-"` // Ignore this field
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Images                []UserProjectImage
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

func (up UserProject) getByIdCat(limit int) []UserProject {
	var ups []UserProject
	db.Limit(limit).Where("is_online = ? AND user_project_category_id = ?", "1", Itoa(int(up.UserProjectCategoryId))).Find(&ups)
	return ups
}

// getAllFromIdUser permet de récupérer tous les projets d'un user
// en fonction de l'id user
func (up UserProject) getAllFromIdUser() []UserProject {
	var ups []UserProject
	db.Where("is_online = ? AND user_id = ?", "1", Itoa(int(up.UserId))).Find(&ups)
	return ups
}

// fonction permettant de récupérer tous les projets actifs
func (up UserProject) getAllProjects() []UserProject {
	var ups []UserProject
	db.Where("is_online = ? ", "1").Find(&ups)
	return ups
}

func (up UserProject) getRandomProjects(numbProjects int) []UserProject {
	ups := up.getAllProjects()
	//numbProjects
	result := make([]UserProject, numbProjects)

	for i := 0; i > numbProjects; i++ {
		result[i] = ups[rand.Intn(len(ups))]
	}

	return result
}
