package controler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	M "github.com/konginteractive/cme/model"
	"log"
	"math"
	. "strconv"
)

var ForumTempl = "forum"
var ForumAddTempl = "forum_add"

var maxElementsInPage = 10

func ForumView() M.Page {

	log.Println("ForumView appelé")

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"
	p.PageLevel = ""
	p.Forums = getListForums()
	p.Categories = getAllFormCategories()
	p.PagesList = createPaginate()

	injectNumPostsToForums(p.Forums)

	return p
}

func ForumAddView() M.Page {

	log.Println("ForumAddView appelé")

	p := new(M.PageForum)
	p.Title = "Titre du sujet"
	p.MainClass = "nouveaututo"
	p.PageLevel = "../"
	p.Forums = make([]M.Forum, 2)
	p.Forums[0].Title = "Test A"
	p.Forums[1].Title = "Test B"

	return p
}

// permet de retourner toutes les catégories
func getAllFormCategories() []M.ForumCategory {
	db := connectToDatabase()
	var cat []M.ForumCategory
	db.Find(&cat)
	return cat
}

func injectNumPostsToForums(forums []M.Forum) []M.Forum {
	lenForum := len(forums)
	for i := 0; i < lenForum; i++ {
		j := forums[i].Id
		forums[i].PostNumb = getNumPostForum(j)
		//log.Print(num)
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

// permet de récupérer le nombre de forums de question total de la base de donnée
func getNumForms() int {
	db := connectToDatabase()
	var forums []M.Forum
	var num int
	db.Where("is_online = ?", "1").Find(&forums).Count(&num)
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
	log.Print(mf)

	p := make([]M.Paginate, nb)

	for i := 0; i < nb; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "/forum/p/" + t
	}
	return p
}

// fonction permetttant de rechercher dans les titres des questions
/*
func searchInTitle(s string) []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Table(&forums).Where("title = ?", s)
	return forums
}
*/
