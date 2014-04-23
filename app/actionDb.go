package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// fonction permettant de se connecter à la base de donnée
func connectToDatabase() gorm.DB {
	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/cme_test?charset=utf8&parseTime=True")
	db.SingularTable(true)
	return db
}

/*
func (db *DB)Close() error{
   var err error
   err = sql.Close()
   return err
}
*/
