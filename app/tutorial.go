package app

import (
	. "strconv"
	"strings"
	"time"
)

type Tutorial struct {
	Id                 int64
	UserId             int64
	TutorialCategoryId int64
	Title              string `sql:"type:varchar(100);"`
	Text               string
	Keywords           string `sql:"type:varchar(200);"`
	IsOnline           int64
	PostNumb           int64  `sql:"-"` // Ignore this field
	CategoryTitle      string `sql:"-"` // Ignore this field
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Posts              []TutorialPost
}

// fonction public
// permet d'enregistrer les éléments du tutoriel
func (t Tutorial) Save() {
	db.Save(&t)
}

// permet de donner l'id d'une question à partir de son titre
func (t Tutorial) getIdFromCatName(cat string) int64 {

	catList := t.getAllCategories()
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		// vérifie si la cat est celle cherchée
		if strings.ToLower(catList[i].Title) == cat {
			return catList[i].Id
		}
	}
	// par défaut retourne la cat 1
	return 1
}

// permet de retourner toutes les catégories
// permet aussi de créer les liens pour les catégories
func (t Tutorial) getAllCategories() []TutorialCategory {

	var cat []TutorialCategory
	db.Find(&cat)

	// create dynamic links
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		//cat[i].Url = "/forum/categorie/" + strings.ToLower(cat[i].Title) + "/" + Itoa(int(cat[i].Id))
		cat[i].Url = "/tutoriel/" + strings.ToLower(cat[i].Title)
	}

	return cat
}

// permet de récupérer toute la listes des tutoriels
// en fonction de la limite affichable par page
func (t Tutorial) getList() []Tutorial {
	var tutorials []Tutorial
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&tutorials)
	return tutorials
}

// permet de récupérer toute la listes des tutoriels
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func (t Tutorial) getListPaged(fromPage int64) []Tutorial {

	var tutorials []Tutorial
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&tutorials)
	return tutorials
}

// permet de récupérer des questions du tutoriel à partir de l'id de la catégorie
func (t Tutorial) getListFromCat(id int64) []Tutorial {

	var tutorials []Tutorial
	db.Limit(maxElementsInPage).Where("is_online = ? and tutorial_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&tutorials)
	return tutorials
}

// permet de récupérer une liste des tutoriels en fonction
// de l'id de la catégorie sélectionnée
// et de la page dans laquelle on est
func (t Tutorial) getListFromCatPaged(id int64, fromPage int64) []Tutorial {

	var tutorials []Tutorial
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ? and tutorial_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&tutorials)
	return tutorials
}

// permet de récupérer le nombre de tutoriels total de la base de donnée
func (t Tutorial) count() int {

	var tutorials []Tutorial
	var num int
	db.Where("is_online = ?", "1").Find(&tutorials).Count(&num)
	return num
}

// permet de récupérer le nombre de tutoriels total de la base de donnée
// en fonction de l'id de la catégorie
func (t Tutorial) countFromIdCat(id int64) int {

	var tutorials []Tutorial
	var num int
	db.Where("is_online = ? and tutorial_category_id = ?", "1", Itoa(int(id))).Find(&tutorials).Count(&num)
	return num
}

// permet de récupérer les posts d'un tutoriel
// à partir de l'id d'un tutoriel
func (t Tutorial) getPost(id int) []TutorialPost {

	idTutorial := Itoa(id)
	var posts []TutorialPost
	db.Where("is_online = ? and tutorial_id = ?", "1", idTutorial).Find(&posts)
	return posts
}

// permet de récupérer le nombre de posts d'un tutoriel
// à partir de l'id d'un tutoriel
func (t Tutorial) countPost(id int64) int64 {

	i := int(id)
	idTutorial := Itoa(i)
	var posts []TutorialPost
	var num int64
	db.Where("is_online = ? and tutorial_id = ?", "1", idTutorial).Find(&posts).Count(&num)
	return num
}

// fonction permetttant de rechercher dans les titres des questions
func (t Tutorial) search(s string) []Tutorial {

	var tutorials []Tutorial
	db.Where("is_online = ? and title LIKE ? ", "1", "%"+s+"%").Or("is_online = ? and text LIKE ? ", "1", "%"+s+"%").Order("id desc").Find(&tutorials)
	return tutorials
}

//
//
//
//
//
//
//
//
//
//
//
// END ♥ allez hop au travail !
