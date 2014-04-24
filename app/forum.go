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
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Posts           []ForumPost
}

/*
// Helper de vue
type ForumViewHelper struct {
	CategoryTitle string
	Forum
}
*/

// fonction public
// permet d'enregistrer les éléments du formulaire
func (f Forum) Save() {
	db.Save(&f)
}

// permet de donner l'id d'une question à partir de son titre
func (f Forum) getIdFromCatName(cat string) int64 {

	catList := f.getAllCategories()
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
func (f Forum) getAllCategories() []ForumCategory {

	var cat []ForumCategory
	db.Find(&cat)

	// create dynamic links
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		//cat[i].Url = "/forum/categorie/" + strings.ToLower(cat[i].Title) + "/" + Itoa(int(cat[i].Id))
		cat[i].Url = "/forum/" + strings.ToLower(cat[i].Title)
	}

	return cat
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
func (f Forum) getListPaged(fromPage int64) []Forum {

	var forums []Forum
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&forums)
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
func (f Forum) getListFromCatPaged(id int64, fromPage int64) []Forum {

	var forums []Forum
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&forums)
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
func (f Forum) getPost(id int) []ForumPost {

	idForum := Itoa(id)
	var posts []ForumPost
	db.Where("is_online = ? and forum_id = ?", "1", idForum).Find(&posts)
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
