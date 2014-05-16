package app

import (
	"errors"
	"log"
	. "strconv"
	"time"
)

type User struct {
	Id         int64
	FirstName  string `sql:"type:varchar(100);"`
	LastName   string `sql:"type:varchar(100);"`
	Text       string
	Email      string
	Pass       string
	Keywords   string
	Facebook   string
	Twitter    string
	LinkedIn   string
	Graduation string
	PhotoProfil string
	IsOnline   int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Images     []UserImage
}

// fonction public
// permet d'enregistrer les éléments du user
// retourne l'id du user
func (u User) Save() int64 {
	db.Save(&u)
	if u.Id != 0 {
		return u.Id
	} else {
		db.Last(&u)
		return u.Id
	}
}

/*
// fonction public
// permet d'enregistrer les éléments du formulaire
func (u User) Save() {
	//log.Println(u)
	db.Save(&u)
}
*/

func (u User) SavePhoto(nouvPhoto string) {
	userId := u.Id
	db.First(&u, userId).Update("PhotoProfil", nouvPhoto)
}

func (u User) SaveBio() {
	userId := u.Id
	userBio := u.Text
	db.First(&u, userId).Update("text", userBio)
}

func (u User) SaveSocial() {
	userId := u.Id
	userFb := u.Facebook
	userTw := u.Twitter
	userLi := u.LinkedIn

	// J'ignore pourquoi mais la ligne suivante ne fonctionne pas
	//db.First(&u, userId).Updates(User{Facebook: userFb, Twitter: userTw, LinkedIn: userLi})

	// Je dois faire trois requêtes différentes pour mettre à jour les infos
	db.First(&u, userId).Update("Facebook", userFb)
	db.First(&u, userId).Update("Twitter", userTw)
	db.First(&u, userId).Update("LinkedIn", userLi)
}

func (u User) SaveGrad() {
	userId := u.Id
	userGrad := u.Graduation
	db.First(&u, userId).Update("graduation", userGrad)
}

func (u User) SaveUserName() {
	userId := u.Id
	userPrenom := u.FirstName
	userNom := u.LastName
	db.First(&u, userId).Updates(User{FirstName: userPrenom, LastName: userNom})
}

func (u User) DeleteAccount() {
	db.Delete(&u)
	// @todo : Déconnecter l'utilisateur en cours et lui présenter la page d'accueil
}

// fonction permettant de retourner un utilisateur
// en fonction de son login et son mot de passe
// retourne un utilisateur si trouvé
// renvoie une erreur le cas échéant
func (u *User) LoginPassExist() (v bool, err error) {

	log.Println("SearchUserLoginAndPass appelé")
	log.Println(u.Email + " / " + u.Pass)

	var user User
	if u.Email != "" && u.Pass != "" {
		db.Where("email = ? AND pass = ?", u.Email, u.Pass).Find(&user)
	} else {
		err = errors.New("Il manque le login et/ou le mot de passe")
	}

	u.Id = user.Id
	u.LastName = user.LastName
	u.FirstName = user.FirstName

	if u.Id != 0 && u.LastName != "" && u.FirstName != "" {
		v = true
	}

	return
}

// permet de récupérer toute les données du user
// en fonction de son id
func (u User) getById() User {
	db.Where("is_online = ? AND id = ?", "1", Itoa(int(u.Id))).Find(&u)
	return u
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func (u User) getList() []User {

	var users []User
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&users)
	return users
}

// permet de récupérer toute la listes des questions du forum
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func (u User) getListPaged(fromPage int64) []User {

	var users []User
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&users)
	return users
}

// permet de récupérer le nombre de forums de question total de la base de donnée
func (u User) count() int {

	var users []User
	var num int
	db.Where("is_online = ?", "1").Find(&users).Count(&num)
	return num
}

// fonction permetttant de rechercher dans les titres des questions
func (u User) search(s string) []User {

	var users []User
	db.Where("is_online = ? and first_name LIKE ? ", "1", "%"+s+"%").Or("is_online = ? and last_name LIKE ? ", "1", "%"+s+"%").Order("id desc").Find(&users)
	return users
}
