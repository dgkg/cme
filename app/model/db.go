package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// fonction permettant de se connecter à la base de donnée
func connectToDatabase() gorm.DB {
	db, _ := gorm.Open("mysql", "root:F4hjsv2REJchspP5@/cme?charset=utf8&parseTime=True")
	db.SingularTable(true)
	return db
}

// connexion à la base de donnée
var db = connectToDatabase()

// donne le nombre d'éléments max à afficher par page
var maxElementsInPage = 3
