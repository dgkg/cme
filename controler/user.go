package controler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	M "github.com/konginteractive/cme/model"
	"log"
)

var UserTempl = "user"

func UserView() M.Page {

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

	log.Println("Étudiants appelé")

	p := new(M.PageUser)
	p.Title = "Élèves"
	p.MainClass = "eleves"

	// pagination
	p.PagesList = make([]M.Paginate, 5)

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

func UserFicheView() M.Page {

	UserTempl = "student_fiche"
	p := new(M.PageUser)
	p.Title = "Nom de l'étudiant sélectionné"
	p.MainClass = "eleve_fiche"

	return p
}
