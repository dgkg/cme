package controler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
)

var UserTempl = "user"

func UserView() Page {

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

	log.Println("Users appelé")

	p := new(PageTutoriels)
	p.Title = "Users"
	p.MainClass = "eleves"

	// pagination
	p.PagesList = make([]Paginate, 5)

	p.PagesList[0].Title = "1"
	p.PagesList[0].Url = "/eleves/page/1"

	p.PagesList[1].Title = "2"
	p.PagesList[1].Url = "/eleves/page/2"

	p.PagesList[2].Title = "3"
	p.PagesList[2].Url = "/eleves/page/3"

	p.PagesList[3].Title = "4"
	p.PagesList[3].Url = "/eleves/page/4"

	p.PagesList[4].Title = "5"
	p.PagesList[4].Url = "/eleves/page/5"

	return p
}