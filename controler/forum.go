package controler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	M "github.com/konginteractive/cme/model"
	"log"
	"math"
	"net/http"
	. "strconv"
	"strings"
)

// template utilisé par les pages affichés
// par défaut le template est form
var ForumTempl = "forum"

// donne le nombre d'éléments max à afficher par page
var maxElementsInPage = 3

// permet d'afficher la liste des questions du forum
func ForumView() M.Page {

	log.Println("ForumView appelé")
	// surcharge de la variable d'affichage
	ForumTempl = "forum"

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"
	p.Forums = forumGetList()
	p.Categories = forumGetAllCategories()
	p.PagesList = forumCreatePaginate()

	forumInjectDataToDisplay(p.Forums)

	return p
}

// permet d'afficher la liste des questions du forum
// avec la fonction de pagination
func ForumViewPaged(page string) M.Page {

	log.Println("ForumViewPaged appelé : " + page)
	// surcharge de la variable d'affichage
	ForumTempl = "forum"

	pagePosition, _ := ParseInt(page, 0, 64)

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"
	p.Forums = forumGetListPaged(pagePosition)
	p.Categories = forumGetAllCategories()
	p.PagesList = forumCreatePaginate()

	forumInjectDataToDisplay(p.Forums)

	return p
}

// permet d'afficher la liste des questions du forum avec la catégorie correspondante
func FormViewCategory(cat string) M.Page {

	log.Println("ForumView appelé")
	// surcharge de la variable d'affichage
	ForumTempl = "forum"

	// récupère l'id de la catégorie
	idCat := forumGetIdFromCatName(cat)

	p := new(M.PageForum)
	p.Title = "Forum " + cat
	p.MainClass = "forum"
	p.Forums = formusGetListFromCat(idCat)
	p.Categories = forumGetAllCategories()
	p.PagesList = forumCreatePaginateFromIdCat(idCat)

	forumInjectDataToDisplay(p.Forums)

	return p
}

// permet d'afficher une liste de questions en fonction de
// la catégorie sélectionnée
// et de la page en cours sélectionnée
func FormViewCategoryPaged(cat string, page string) M.Page {

	log.Println("ForumView appelé")
	// surcharge de la variable d'affichage
	ForumTempl = "forum"

	pagePosition, _ := ParseInt(page, 0, 64)
	// récupère l'id de la catégorie
	idCat := forumGetIdFromCatName(cat)

	p := new(M.PageForum)
	p.Title = "Forum " + cat
	p.MainClass = "forum"
	p.Forums = formusGetListFromCatPaged(idCat, pagePosition)
	p.Categories = forumGetAllCategories()
	p.PagesList = forumCreatePaginateFromIdCat(idCat)

	forumInjectDataToDisplay(p.Forums)

	return p
}

// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les titres
// et va chercher dans le texte du titre
func ForumViewSearch(q string) M.Page {

	log.Println("ForumViewSearch appelé")
	// surcharge de la variable d'affichage
	ForumTempl = "forum"

	p := new(M.PageForum)
	p.Title = "Forum Rechercher"
	p.MainClass = "forum"
	p.Forums = formSearch(q)
	p.Categories = forumGetAllCategories()

	if len(p.Forums) == 0 {
		p.SearchText = q
	}

	forumInjectDataToDisplay(p.Forums)

	return p
}

// permet d'afficher le formulaire de création d'une question du formulaire
func ForumViewAdd() M.Page {

	log.Println("ForumAddView appelé")
	// surcharge de la variable d'affichage
	ForumTempl = "forum_add"

	p := new(M.PageForum)
	p.Title = "Titre du sujet"
	p.MainClass = "nouveausujet"
	p.Forums = make([]M.Forum, 2)
	p.Categories = forumGetAllCategories()

	return p
}

// permet de donner l'id d'une question à partir de son titre
func forumGetIdFromCatName(cat string) int64 {
	catList := forumGetAllCategories()
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
func forumGetAllCategories() []M.ForumCategory {
	db := connectToDatabase()
	var cat []M.ForumCategory
	db.Find(&cat)

	// create dynamic links
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		//cat[i].Url = "/forum/categorie/" + strings.ToLower(cat[i].Title) + "/" + Itoa(int(cat[i].Id))
		cat[i].Url = "/forum/" + strings.ToLower(cat[i].Title)
	}

	return cat
}

// Permet de retrouver le nombre de réponses pour chaque post
// Permet aussi de réduire la description du texte de desc à 250 caractères
func forumInjectDataToDisplay(forums []M.Forum) []M.Forum {
	lenForum := len(forums)

	for i := 0; i < lenForum; i++ {
		id := forums[i].Id
		// permet de réaliser des extraits si le texte est trop long
		if len(forums[i].Text) > 250 {
			text := forums[i].Text[0:250]
			forums[i].Text = text
		}
		forums[i].PostNumb = forumCountPost(id)
	}

	return forums
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func forumGetList() []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer toute la listes des questions du forum
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func forumGetListPaged(fromPage int64) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer des questions du forum à partir de l'id de la catégorie
func formusGetListFromCat(id int64) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer une liste de questions du formulaire en fonction
// de l'id de la catégorie sélectionnée
// et de la page dans laquelle on est
func formusGetListFromCatPaged(id int64, fromPage int64) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&forums)
	return forums
}

// permet de récupérer le nombre de forums de question total de la base de donnée
func forumCount() int {
	db := connectToDatabase()
	var forums []M.Forum
	var num int
	db.Where("is_online = ?", "1").Find(&forums).Count(&num)
	return num
}

// permet de récupérer le nombre de forums de question total de la base de donnée
// en fonction de l'id de la catégorie
func forumCountFromIdCat(id int64) int {
	db := connectToDatabase()
	var forums []M.Forum
	var num int
	db.Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Find(&forums).Count(&num)
	return num
}

// permet de récupérer les posts d'un forum
// à partir de l'id d'un forum
func forumGetPost(id int) []M.ForumPost {
	idForum := Itoa(id)
	db := connectToDatabase()
	var posts []M.ForumPost
	db.Where("is_online = ? and forum_id = ?", "1", idForum).Find(&posts)
	return posts
}

// permet de récupérer le nombre de posts d'un forum
// à partir de l'id d'un forum
func forumCountPost(id int64) int64 {
	i := int(id)
	idForum := Itoa(i)
	db := connectToDatabase()
	var posts []M.ForumPost
	var num int64
	db.Where("is_online = ? and forum_id = ?", "1", idForum).Find(&posts).Count(&num)
	return num
}

// fonction permettant de se connecter à la base de donnée
func connectToDatabase() gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/cme_test?charset=utf8&parseTime=True")
	db.SingularTable(true)
	return db
}

// fonction pour créer la pagination
func forumCreatePaginate() []M.Paginate {
	elTotal := forumCount()

	nb := elTotal / maxElementsInPage
	mf := int(math.Floor(float64(nb)))
	p := make([]M.Paginate, nb)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
}

// fonction pour créer la pagination à partir d'une catégorie sélectionnée
func forumCreatePaginateFromIdCat(id int64) []M.Paginate {
	elTotal := forumCountFromIdCat(id)

	nb := elTotal / maxElementsInPage
	mf := int(math.Floor(float64(nb)))
	p := make([]M.Paginate, nb)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
}

// permet de valider les éléments du formulaire
func ForumValidateForm(r *http.Request) bool {

	if r.PostFormValue("post-nom") == "" {
		return false
	}

	if r.PostFormValue("post-contenu") == "" {
		return false
	}

	return true
}

// permet d'enregistrer les éléments du formulaire
func ForumSaveForm(r *http.Request) {
	db := connectToDatabase()
	var f M.Forum
	f.Title = r.PostFormValue("post-nom")
	f.Text = r.PostFormValue("post-contenu")
	f.ForumCategoryId, _ = ParseInt(r.PostFormValue("post-cat"), 0, 64)
	f.IsOnline = 1 // permet de mettre en ligne automatiquement la quesion
	db.Save(&f)
}

// fonction permetttant de rechercher dans les titres des questions
func formSearch(s string) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
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
