package app

import (
	. "strconv"
	"strings"
	"time"
)

type Forum struct {
	Id              int64
	UserId          int64
	ForumCategoryId int64
	Title           string `sql:"type:varchar(100);"`
	Text            string
	Keywords        string `sql:"type:varchar(200);"`
	IsSolved        int64
	IsOnline        int64
	PostNumb        int64  `sql:"-"` // Ignore this field
	CategoryTitle   string `sql:"-"` // Ignore this field
	Url             string `sql:"-"` // Ignore this field
	UserName        string `sql:"-"` // Ignore this field
	CreatedAtString string `sql:"-"` // Ignore this field
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Posts           []ForumPost
}

// fonction public
// permet d'enregistrer les éléments du formulaire
func (f Forum) Save() int64 {
	db.Save(&f)
	if f.Id != 0 {
		return f.Id
	} else {
		db.Last(&f)
		return f.Id
	}
}

// fonction public
// permet d'enregistrer les modifications des éléments du formulaire
func (f Forum) Update() int64 {
	db.First(&f, &f.Id).Update(&f)
	//return f.Id
	return 30
}

// fonction public
// permet de supprimer une question du forum
func (f Forum) Delete() {
	db.Delete(&f)
}

// permet de donner l'id d'une question à partir de son titre
func (f Forum) getIdFromCatName(cat string) int64 {

	catList := f.getAllCategories(0)
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
func (f Forum) getAllCategories(id int64) []ForumCategory {

	var cat []ForumCategory
	db.Find(&cat)

	// create dynamic links
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		// permet de créer un lien url
		cat[i].Url = "/forum/" + strings.ToLower(cat[i].Title)
		// permet de créer la sélection d'un élément
		if cat[i].Id == id {
			cat[i].IsSelected = true
		} else {
			cat[i].IsSelected = false
		}
	}
	return cat
}

// permet de récupérer toute la question du forum
// en fonction de son id
func (f Forum) getById() Forum {
	db.Where("is_online = ? AND id = ?", "1", Itoa(int(f.Id))).Find(&f)
	return f
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func (f Forum) getList() []Forum {

	var forums []Forum
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer toute la listes des questions du forum
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func (f Forum) getListPaged(fromPage int) []Forum {
	//
	var forums []Forum
	p := fromPage * maxElementsInPage
	p = p - maxElementsInPage
	db.Limit(maxElementsInPage).Offset(p).Where("is_online = ?", "1").Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer des questions du forum à partir de l'id de la catégorie
func (f Forum) getListFromCat(id int64) []Forum {

	var forums []Forum
	db.Limit(maxElementsInPage).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer une liste de questions du formulaire en fonction
// de l'id de la catégorie sélectionnée
// et de la page dans laquelle on est
func (f Forum) getListFromCatPaged(id int64, fromPage int) []Forum {

	var forums []Forum
	p := fromPage * maxElementsInPage
	p = p - maxElementsInPage
	db.Limit(maxElementsInPage).Offset(p).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer le nombre de forums de question total de la base de donnée
func (f Forum) count() int {

	var forums []Forum
	var num int
	db.Where("is_online = ?", "1").Find(&forums).Count(&num)
	return num
}

// permet de récupérer le nombre de forums de question total de la base de donnée
// en fonction de l'id de la catégorie
func (f Forum) countFromIdCat(id int64) int {

	var forums []Forum
	var num int
	db.Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Find(&forums).Count(&num)
	return num
}

// permet de récupérer les posts d'un forum
// à partir de l'id d'un forum
func (f Forum) getPost() []ForumPost {
	var posts []ForumPost
	db.Where("is_online = ? and forum_id = ?", "1", Itoa(int(f.Id))).Find(&posts)
	return posts
}

// permet de récupérer le nombre de posts d'un forum
// à partir de l'id d'un forum
func (f Forum) countPost(id int64) int64 {

	i := int(id)
	idForum := Itoa(i)
	var posts []ForumPost
	var num int64
	db.Where("is_online = ? and forum_id = ?", "1", idForum).Find(&posts).Count(&num)
	return num
}

// fonction permetttant de rechercher dans les titres des questions
func (f Forum) search(s string) []Forum {

	var forums []Forum
	db.Where("is_online = ? and title LIKE ? ", "1", "%"+s+"%").Or("is_online = ? and text LIKE ? ", "1", "%"+s+"%").Order("id desc").Find(&forums)
	return forums
}
