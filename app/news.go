package app

import (
	"time"
)

type News struct {
	Id             int64
	UserId         int64
	UserName	   string `sql:"-"` // Ignore this field
	NewsCategoryId int64
	Title          string `sql:"type:varchar(100);"`
	Text           string
	Keywords       string `sql:"type:varchar(200);"`
	IsOnline       int64
	CreatedAt      time.Time
	CreatedAtString	string `sql:"-"` // Ignore this field
	UpdatedAt      time.Time
	Posts          []NewsPost
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func (n News) getList() []News {

	var news []News
	db.Limit(2).Where("is_online = ?", "1").Order("id desc").Find(&news)
	return news
}
