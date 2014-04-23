package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var TutorialTempl = "tutorial"
var TutorialAddTempl = "tutorial_add"

func (pt PageTutoriels) View() Page {

	// surcharge de la variable d'affichage
	Templ = "tutorial"

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

	pt.Title = "Tutoriels"
	pt.MainClass = "tutoriels"

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
	//p.Forums = make([]M.Forum, 2)
	//p.Forums[0].Title = "Test A"
	//p.Forums[1].Title = "Test B"

	return pt
}
