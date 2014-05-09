package app

import (
	"strings"
	"time"
)

type News struct {
	Id              int64
	UserId          int64
	UserName        string `sql:"-"` // Ignore this field
	NewsCategoryId  int64
	Title           string `sql:"type:varchar(100);"`
	Text            string
	Keywords        string `sql:"type:varchar(200);"`
	IsOnline        int64
	CreatedAt       time.Time
	CreatedAtString string `sql:"-"` // Ignore this field
	UpdatedAt       time.Time
	Posts           []NewsPost
}

// fonction public
// permet d'enregistrer les éléments du formulaire
func (n News) Save() {
	db.Save(&n)
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func (n News) getList(numbElements int) []News {

	var news []News
	db.Limit(numbElements).Where("is_online = ?", "1").Order("id desc").Find(&news)
	return news
}

// permet de récupérer toute la listes des news
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func (n News) getListPaged(fromPage int64) []News {

	var news []News
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&news)
	return news
}

// permet de retourner toutes les catégories
// permet aussi de créer les liens pour les catégories
func (n News) getAllCategories() []NewsCategory {

	var cat []NewsCategory
	db.Find(&cat)

	// create dynamic links
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		cat[i].Url = "/news/" + strings.ToLower(cat[i].Title)
	}

	return cat
}

// permet de récupérer le nombre de news total de la base de donnée
func (n News) count() int {

	var news []News
	var num int
	db.Where("is_online = ?", "1").Find(&news).Count(&num)
	return num
}
