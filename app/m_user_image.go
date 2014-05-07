package app

import (
	//"fmt"
	"time"
)

type UserImage struct {
	Id                  int64
	UserId              int64
	UserImageCategoryId int64
	Title               string `sql:"type:varchar(100);"`
	Url                 string `sql:"type:varchar(200);"`
	Description         string
	DescriptionCourte   string
	IsFeatured          int64
	IsOnline            int64
	Year				int64
	Month               int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (u UserImage) getProjets(nbProjets int) []UserImage {

	// Récupérer les images
	var images []UserImage
	db.Limit(nbProjets).Where("is_online = ?", "1").Order("id desc").Find(&images)

	return images
}
