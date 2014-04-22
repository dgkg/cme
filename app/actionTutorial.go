package controler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	M "github.com/konginteractive/cme/model"
	"log"
)

var TutorialTempl = "tutorial"
var TutorialAddTempl = "tutorial_add"

func TutorialView() Page {

	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/cme_test?charset=utf8&parseTime=True")
	//cme_test
	db.SingularTable(true)

	//db.Find(&forum)

	/////////////
	/////////////
	/////////////
	/////////////
	/////////////
	/////////////

	log.Println("Tutorials appelé")

	p := new(PageTutoriels)
	p.Title = "Tutoriels"
	p.MainClass = "tutoriels"

	// Création des forums
	/*
		p.News = make([]M.NewsViewHelper, 2)

		p.News[0].Id = 1
		p.News[0].Title = "Test A"
		p.News[0].CategoryTitle = "Packaging"

		p.News[1].Id = 2
		p.News[1].Title = "Test B"
		p.News[1].CategoryTitle = "Logiciel"
	*/

	// pagination
	p.PagesList = make([]Paginate, 5)

	p.PagesList[0].Title = "1"
	p.PagesList[0].Url = "/tutoriels/page/1"

	p.PagesList[1].Title = "2"
	p.PagesList[1].Url = "/tutoriels/page/2"

	p.PagesList[2].Title = "3"
	p.PagesList[2].Url = "/tutoriels/page/3"

	p.PagesList[3].Title = "4"
	p.PagesList[3].Url = "/tutoriels/page/4"

	p.PagesList[4].Title = "5"
	p.PagesList[4].Url = "/tutoriels/page/5"

	return p
}

func TutorialAddView() Page {

	log.Println("TutoAddView appelé")

	p := new(PageTutoriels)
	p.Title = "Titre du tuto"
	p.MainClass = "nouveaututo"
	//p.Forums = make([]M.Forum, 2)
	//p.Forums[0].Title = "Test A"
	//p.Forums[1].Title = "Test B"

	return p
}