package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
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
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Posts              []TutorialPost
}

type TutorialCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TutorialPost struct {
	Id         int64
	TutorialId int64
	UserId     int64
	Text       string
	IsOline    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TutorialsViewHelper struct {
	Tutorial
}

type PageTutoriels struct {
	PagesList []Paginate
	Tutoriels []TutorialsViewHelper
	PageWeb
}

var TutorialTempl = "tutorial"
var TutorialAddTempl = "tutorial_add"

func (pt PageTutoriels) View() Page {

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/cme_test?charset=utf8&parseTime=True")
	//cme_test
	db.SingularTable(true)

	log.Println("Tutorials appelé")

	pt.Title = "Tutoriels"
	pt.MainClass = "tutoriels"

	// pagination
	pt.PagesList = make([]Paginate, 5)

	pt.PagesList[0].Title = "1"
	pt.PagesList[0].Url = "/tutoriels/page/1"

	pt.PagesList[1].Title = "2"
	pt.PagesList[1].Url = "/tutoriels/page/2"

	pt.PagesList[2].Title = "3"
	pt.PagesList[2].Url = "/tutoriels/page/3"

	pt.PagesList[3].Title = "4"
	pt.PagesList[3].Url = "/tutoriels/page/4"

	pt.PagesList[4].Title = "5"
	pt.PagesList[4].Url = "/tutoriels/page/5"

	return pt
}

func (pt PageTutoriels) AddView() Page {

	log.Println("TutoAddView appelé")

	// surcharge de la variable d'affichage
	Templ = "tutorial_add"

	pt.Title = "Titre du tuto"
	pt.MainClass = "nouveaututo"

	return pt
}
