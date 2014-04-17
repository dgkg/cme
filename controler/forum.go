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

var ForumTempl = "forum"
var ForumAddTempl = "forum_add"

var maxElementsInPage = 3

// permet d'afficher la liste des questions du forum
func ForumView() M.Page {

	log.Println("ForumView appelé")

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"
	p.PageLevel = ""
	p.Forums = getListForums()
	p.Categories = getAllFormCategories()
	p.PagesList = createPaginate()

	injectDataForumToDisplay(p.Forums)

	return p
}

// permet d'afficher la liste des questions du forum
func ForumViewPaged(page string) M.Page {

	log.Println("ForumViewPaged appelé : " + page)
	pagePosition, _ := ParseInt(page, 0, 64)

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"
	p.PageLevel = ""
	p.Forums = getListForumsPaged(pagePosition)
	p.Categories = getAllFormCategories()
	p.PagesList = createPaginate()

	injectDataForumToDisplay(p.Forums)

	return p
}

// permet d'afficher la liste des questions du forum avec la catégorie correspondante
func FormViewCategory(cat string) M.Page {
	log.Println("ForumView appelé")

	// récupère l'id de la catégorie
	idCat := getIdFromCatName(cat)

	p := new(M.PageForum)
	p.Title = "Forum " + cat
	p.MainClass = "forum"
	p.PageLevel = "../"
	p.Forums = getListFormusFromCat(idCat)
	p.Categories = getAllFormCategories()
	p.PagesList = createPaginateFromIdCat(idCat)

	injectDataForumToDisplay(p.Forums)

	return p
}

func FormViewCategoryPaged(cat string, page string) M.Page {

	log.Println("ForumView appelé")
	pagePosition, _ := ParseInt(page, 0, 64)
	// récupère l'id de la catégorie
	idCat := getIdFromCatName(cat)

	p := new(M.PageForum)
	p.Title = "Forum " + cat
	p.MainClass = "forum"
	p.PageLevel = "../"
	p.Forums = getListFormusFromCatPaged(idCat, pagePosition)
	p.Categories = getAllFormCategories()
	p.PagesList = createPaginateFromIdCat(idCat)

	injectDataForumToDisplay(p.Forums)

	return p
}

// permet d'afficher le formulaire de création d'une question du formulaire
func ForumAddView() M.Page {

	log.Println("ForumAddView appelé")

	p := new(M.PageForum)
	p.Title = "Titre du sujet"
	p.MainClass = "nouveausujet"
	p.PageLevel = "../"
	p.Forums = make([]M.Forum, 2)
	p.Categories = getAllFormCategories()

	return p
}

// permet de donner l'id d'une question à partir de son titre
func getIdFromCatName(cat string) int64 {
	catList := getAllFormCategories()
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
func getAllFormCategories() []M.ForumCategory {
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
func injectDataForumToDisplay(forums []M.Forum) []M.Forum {
	lenForum := len(forums)

	for i := 0; i < lenForum; i++ {
		id := forums[i].Id
		text := forums[i].Text[0:250]
		forums[i].PostNumb = getNumPostForum(id)
		forums[i].Text = text
	}

	return forums
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func getListForums() []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Find(&forums)
	return forums
}

// permet de récupérer toute la listes des questions du forum
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func getListForumsPaged(fromPage int64) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Find(&forums)
	return forums
}

// permet de récupérer des questions du forum à partir de l'id de la catégorie
func getListFormusFromCat(id int64) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Find(&forums)
	return forums
}

func getListFormusFromCatPaged(id int64, fromPage int64) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Find(&forums)
	return forums
}

// permet de récupérer le nombre de forums de question total de la base de donnée
func getNumForms() int {
	db := connectToDatabase()
	var forums []M.Forum
	var num int
	db.Where("is_online = ?", "1").Find(&forums).Count(&num)
	return num
}

// permet de récupérer le nombre de forums de question total de la base de donnée
// en fonction de l'id de la catégorie
func getNumFormsFromIdCat(id int64) int {
	db := connectToDatabase()
	var forums []M.Forum
	var num int
	db.Where("is_online = ? and forum_category_id = ?", "1", Itoa(int(id))).Find(&forums).Count(&num)
	return num
}

// permet de récupérer les posts d'un forum
// à partir de l'id d'un forum
func getPostForum(id int) []M.ForumPost {
	idForum := Itoa(id)
	db := connectToDatabase()
	var posts []M.ForumPost
	db.Where("is_online = ? and forum_id = ?", "1", idForum).Find(&posts)
	return posts
}

// permet de récupérer le nombre de posts d'un forum
// à partir de l'id d'un forum
func getNumPostForum(id int64) int64 {
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
func createPaginate() []M.Paginate {
	elTotal := getNumForms()

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
func createPaginateFromIdCat(id int64) []M.Paginate {
	elTotal := getNumFormsFromIdCat(id)

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

func ValidateForum(r *http.Request) bool {

	if r.PostFormValue("post-nom") == "" {
		return false
	}

	if r.PostFormValue("post-contenu") == "" {
		return false
	}

	return true
}

func SetPostForum(r *http.Request) {
	db := connectToDatabase()
	var f M.Forum
	f.Title = r.PostFormValue("post-nom")
	f.Text = r.PostFormValue("post-contenu")
	//f.ForumCategoryId = ParseInt(r.PostFormValue("post-cat"), 0, 64)
	//log.Print(len(r.PostFormValue("post-cat")))
	postCategorie := r.PostFormValue("post-cat")
	log.Println(postCategorie)
	/* for key, _ := range r.PostFormValue("post-cat") {
		//log.Println(Itoa(int(r.PostFormValue("post-cat")[key])))
		log.Println(Itoa(int(r.MultipartForm("post-cat")[key])))
	} */
	//log.Println(r.PostFormValue("post-cat")["value"])
	db.Save(&f)
}

/*
// fonction permetttant de rechercher dans les titres des questions
func searchInTitle(s string) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Table(&forums).Where("title = ?", s)
	return forums
}
*/
