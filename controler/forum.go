package controler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	M "github.com/konginteractive/cme/model"
	"log"
)

var ForumTempl = "forum"

func ForumView() M.Page {

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

	log.Println("Forum appelé")

	p := new(M.PageForum)
	p.Title = "Forum"
	p.MainClass = "forum"

	// Création des forums
	p.Forums = make([]M.ForumViewHelper, 2)

	p.Forums[0].Id = 1
	p.Forums[0].Title = "Test A"
	p.Forums[0].CategoryTitle = "Packaging"

	p.Forums[1].Id = 2
	p.Forums[1].Title = "Test B"
	p.Forums[1].CategoryTitle = "Logiciel"

	// pagination
	p.PagesList = make([]M.Paginate, 5)

	p.PagesList[0].Title = "1"
	p.PagesList[0].Url = "/forum/page/1"

	p.PagesList[1].Title = "2"
	p.PagesList[1].Url = "/forum/page/2"

	p.PagesList[2].Title = "3"
	p.PagesList[2].Url = "/forum/page/3"

	p.PagesList[3].Title = "4"
	p.PagesList[3].Url = "/forum/page/4"

	p.PagesList[4].Title = "5"
	p.PagesList[4].Url = "/forum/page/5"

	return p
}

var ForumAddTempl = "forum_add"

func ForumAddView() M.Page {
	p := new(M.PageForum)
	p.Title = "Titre du sujet"

	p.Forums = make([]M.ForumViewHelper, 2)

	p.Forums[0].Title = "Test A"
	p.Forums[1].Title = "Test B"

	return p
}

func initDb() {

}
