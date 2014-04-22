package model

import (
	"time"
)

type User struct {
	Id        int64
	FirstName string `sql:"type:varchar(100);"`
	LastName  string `sql:"type:varchar(100);"`
	UrlImage  string // Photo profil de l'étudiant
	Text      string // Biographie
	Email     string
	Pass      string
	Keywords  string //
	IsOnline  int64  // Si la fiche est en ligne
	CreatedAt time.Time
	UpdatedAt time.Time
	Images    []UserImage // Tableau contenu les projets
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
