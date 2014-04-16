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

var maxElementsInPage = 5

func ForumView() M.Page {

	log.Println("ForumView appelé")

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"
	p.PageLevel = ""
	p.Forums = getListForums()
	p.Categories = getAllFormCategories()
	p.PagesList = createPaginate()

	return p
}

func ForumAddView() M.Page {

	log.Println("ForumAddView appelé")

	p := new(M.PageForum)
	p.Title = "Titre du sujet"
	p.Forums = make([]M.Forum, 2)
	p.Forums[0].Title = "Test A"
	p.Forums[1].Title = "Test B"

	ForumTempl = "forum_add"

	return p
}

// permet de retourner toutes les catégories
func getAllFormCategories() []M.ForumCategory {
	db := connectToDatabase()
	var cat []M.ForumCategory
	db.Find(&cat)
	return cat
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func getListForums() []M.Forum {
	db := connectToDatabase()
	var forums []M.Forum
	db.Limit(maxElementsInPage).Find(&forums)
	return forums
}

// permet de récupérer le nombre de forums de question total de la base de donnée
func getNumForms() int {
	db := connectToDatabase()
	var forums []M.Forum
	var num int
	db.Find(&forums).Count(&num)
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

<<<<<<< HEAD
func ForumAddView() M.Page {
	p := new(M.PageForum)
	p.Title = "Titre du sujet"
	p.PageLevel = "../"
=======
	nb := elTotal / maxElementsInPage
>>>>>>> 2fef960b0f38d40039e20964eae64f4c667b523d

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

func searchInTitle() {
	/*
		//Work's
		rows := db.Table("forum").Where("title = ?", "Titre B").Select("title, text").Row()

		var title string
		var text string
		rows.Scan(&title, &text)
		log.Print(title + " // " + text)
	*/
}
