package app

import (
	//"fmt"
	. "strconv"
	"strings"
	"time"
)

type UserProjectImage struct {
	Id                  int64
	ProjectId           int64
	UserImageCategoryId int64
	Title               string `sql:"type:varchar(100);"`
	Url                 string `sql:"type:varchar(200);"`
	Description         string
	DescriptionCourte   string
	IsFeatured          int64
	IsOnline            int64
	Year                int64
	Month               int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (ui UserProjectImage) getProjets(nbProjets int) []UserProjectImage {

	// Récupérer les images
	var images []UserProjectImage
	db.Limit(nbProjets).Where("is_online = ?", "1").Order("id desc").Find(&images)

	return images
}

/**********************************************/

// fonction public
// permet d'enregistrer les éléments du formulaire
func (ui UserProjectImage) Save() {
	db.Save(&ui)
}

// getAllByIdProject permet de retourner la liste des Images
// en fonction de l'id du projet
func (ui *UserProjectImage) getAllByIdProject() []UserProjectImage {
	var uis []UserProjectImage
	db.Where("is_online = ? AND id = ?", "1", Itoa(int(ui.ProjectId))).Find(&uis)
	return uis
}

// permet de donner l'id d'une question à partir de son titre
func (ui UserProjectImage) getIdFromCatName(cat string) int64 {

	catList := ui.getAllCategories()
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		// vérifie si la cat est celle cherchée
		if strings.ToLower(catList[i].Title) == cat {
			return catList[i].Id
		}
	}
	// par défaut retourne la cat 1
	return 1
}

// permet de retourner toutes les catégories
// permet aussi de créer les liens pour les catégories
func (ui UserProjectImage) getAllCategories() []UserProjectCategory {

	var cat []UserProjectCategory
	db.Find(&cat)

	// create dynamic links
	lenCat := len(cat)
	for i := 0; i < lenCat; i++ {
		//cat[i].Url = "/forum/categorie/" + strings.ToLower(cat[i].Title) + "/" + Itoa(int(cat[i].Id))
		cat[i].Url = "/images/" + strings.ToLower(cat[i].Title)
	}

	return cat
}

// permet de récupérer toute la question du forum
// en fonction de son id
func (ui UserProjectImage) getById() UserProjectImage {
	db.Where("is_online = ? AND id = ?", "1", Itoa(int(ui.Id))).Find(&ui)
	return ui
}

// permet de récupérer toute la listes des questions du forum
// en fonction de la limite affichable par page
func (ui UserProjectImage) getList() []UserProjectImage {

	var userImages []UserProjectImage
	db.Limit(maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&userImages)
	return userImages
}

// permet de récupérer toute la listes des questions du forum
// avec les fonctions de pagination
// en fonction de la limite affichable par page
func (ui UserProjectImage) getListPaged(fromPage int64) []UserProjectImage {

	var userImages []UserProjectImage
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ?", "1").Order("id desc").Find(&userImages)
	return userImages
}

// permet de récupérer des questions du forum à partir de l'id de la catégorie
func (ui UserProjectImage) getListFromCat(id int64) []UserProjectImage {

	var userImages []UserProjectImage
	db.Limit(maxElementsInPage).Where("is_online = ? and user_image_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&userImages)
	return userImages
}

// permet de récupérer une liste de questions du formulaire en fonction
// de l'id de la catégorie sélectionnée
// et de la page dans laquelle on est
func (ui UserProjectImage) getListFromCatPaged(id int64, fromPage int64) []UserProjectImage {

	var userImages []UserProjectImage
	db.Limit(maxElementsInPage).Offset(int(fromPage)*maxElementsInPage).Where("is_online = ? and user_image_category_id = ?", "1", Itoa(int(id))).Order("id desc").Find(&userImages)
	return userImages
}

// permet de récupérer le nombre de forums de question total de la base de donnée
func (ui UserProjectImage) count() int {

	var userImages []UserProjectImage
	var num int
	db.Where("is_online = ?", "1").Find(&userImages).Count(&num)
	return num
}

// permet de récupérer le nombre de forums de question total de la base de donnée
// en fonction de l'id de la catégorie
func (ui UserProjectImage) countFromIdCat(id int64) int {

	var userImages []UserProjectImage
	var num int
	db.Where("is_online = ? and user_image_category_id = ?", "1", Itoa(int(id))).Find(&userImages).Count(&num)
	return num
}

// fonction permetttant de rechercher dans les titres des questions
func (ui UserProjectImage) search(s string) []UserProjectImage {

	var userImages []UserProjectImage
	db.Where("is_online = ? and title LIKE ? ", "1", "%"+s+"%").Order("id desc").Find(&userImages)
	return userImages
}
