package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

/*
type UsersViewHelper struct {
	User
}
*/

type PageUser struct {
	PagesList []Paginate
	Users     []UsersViewHelper
	PageWeb
}

func (pu PageUser) View() Page {

	// surcharge de la variable d'affichage
	Templ = "user"

	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/cme_test?charset=utf8&parseTime=True")
	//cme_test
	db.SingularTable(true)

	log.Println("Users appelé")

	pu.Title = "Users"
	pu.MainClass = "eleves"

	// pagination
	pu.PagesList = make([]Paginate, 5)

	pu.PagesList[0].Title = "1"
	pu.PagesList[0].Url = "/eleves/page/1"

	pu.PagesList[1].Title = "2"
	pu.PagesList[1].Url = "/eleves/page/2"

	pu.PagesList[2].Title = "3"
	pu.PagesList[2].Url = "/eleves/page/3"

	pu.PagesList[3].Title = "4"
	pu.PagesList[3].Url = "/eleves/page/4"

	pu.PagesList[4].Title = "5"
	pu.PagesList[4].Url = "/eleves/page/5"

	return pu
}
