package app

import (
	"time"
)

type News struct {
	Id             int64
	UserId         int64
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
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&news)
	return news
}
