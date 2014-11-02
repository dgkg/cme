package model

import (
	"log"
	//"reflect"
	. "strconv"
	"time"
)

type UserProjectCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	Url       string `sql:"_"` // Ignore this field
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserProjectCategoryResult struct {
	Id int64
}

// Save permet d'enregistrer les éléments du formulaire
// si l'élément n'existe pas l'ajoute dans la base de donnée
// si l'élément exite met à jour dans la base de donnée
func (upc UserProjectCategory) Save() int64 {
	db.Save(&upc)
	if upc.Id != 0 {
		return upc.Id
	} else {
		db.Last(&upc)
		return upc.Id
	}
}

// Delete permet de supprimer une catégorie d'une image
func (upc UserProjectCategory) Delete() {
	db.Delete(&upc)
}

// getById permet de récupérer toute les données d'une catégorie
// en fonction de son id
func (upc UserProjectCategory) GetById() UserProjectCategory {
	db.Where("id = ?", Itoa(int(upc.Id))).Find(&upc)
	return upc
}

// fonction permettant de récupérer tous les catégories
func (upc UserProjectCategory) GetAllCategories() []UserProjectCategory {
	var ups []UserProjectCategory
	db.Find(&ups)
	return ups
}

// permet de retourner la liste des catégories dans laquelle il y a des projets
func (upc UserProjectCategory) GetCategoriesProjects() []UserProjectCategory {
	//
	var rs []UserProjectCategoryResult
	db.Table("user_project").Select("user_project_category_id as id").Group("user_project_category_id").Scan(&rs)
	//
	upcs := make([]UserProjectCategory, len(rs))
	//
	for key, _ := range rs {
		upc.Id = rs[key].Id
		upc = upc.GetById()
		log.Println(upc.Title)
		upcs[key] = upc
	}
	return upcs
}
