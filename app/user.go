package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type PageUser struct {
	PagesList []Paginate
	Users     []UsersViewHelper
	PageWeb
}

type User struct {
	Id        int64
	FirstName string `sql:"type:varchar(100);"`
	LastName  string `sql:"type:varchar(100);"`
	Text      string
	Email     string
	Pass      string
	Keywords  string
	IsOnline  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Images    []UserImage
}

type UserImage struct {
	Id                  int64
	UserId              int64
	UserImageCategoryId int64
	Title               string `sql:"type:varchar(100);"`
	Url                 string `sql:"type:varchar(200);"`
	Description         string
	IsFeatured          int64
	IsOnline            int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type UserImageCategory struct {
	Id        int64
	Title     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersViewHelper struct {
	User
}

var UserTempl = "user"

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
