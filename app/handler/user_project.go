package handler

import (
	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/helper"
	. "github.com/konginteractive/cme/app/model"
	"net/http"
	. "strconv"
)

// Public function
// permet de récupérer une liste de projets
func UserProjectAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	//page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 {
		// envoie un message d'erreur
		renderString(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.GetByIdCat(3)
		// encodage en json des projets
		renderJson(w, ups)
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
func UserProjectMoreAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if page == 0 {
		// envoie un message d'erreur
		renderString(w, "error")
	} else {

		var up UserProject
		ups := up.GetListPaged(page, 3)
		// encodage en json des projets
		renderJson(w, ups)
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
// et avec une catégorie sélectionnée
func UserProjectMoreInCatAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 || page == 0 {
		// envoie un message d'erreur
		renderString(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.GetByIdCatPaged(page, 3)
		// encodage en json des projets
		renderJson(w, ups)
	}
}

// permet de créer un userProject à partir d'un formulaire HTML
// et permet d'uploader une image associée
// et permet de croper l'image en fonction de sa vignette
func UPCreateAjaxHandler(w http.ResponseWriter, r *http.Request) {

	// création d'un user project
	var up UserProject
	// récupère l'id du user
	UserId, err := ParseInt(r.PostFormValue("id_user"), 0, 64)
	if err != nil {
		// envoie un message d'erreur
		renderString(w, "error")
	}
	up.UserId = UserId
	// crécuprère l'id de la cat
	up.UserProjectCategoryId, err = ParseInt(r.PostFormValue("id_cat"), 0, 64)
	if err != nil {
		// envoie un message d'erreur
		renderString(w, "error")
	}
	// récupère le titre
	up.Title = sanitize.HTML(r.PostFormValue("title"))
	// récupère la description
	up.Description = sanitize.HTML(r.PostFormValue("Description"))
	// permet d'uploader l'image
	up.Url, err = UploadImage(URL_PROJECT_IMAGES, r)
	if err != nil {
		// envoie un message d'erreur
		renderString(w, "error")
	}
	// permet de cropper l'image qui viens d'être uploadée
	err = CropImage(up.Url, 300, 300)
	if err != nil {
		// envoie un message d'erreur
		renderString(w, "error")
	}
	// finit de créer l'objet
	// définit le projet non visible
	up.IsOnline = 0
	// sauvegarde dans la base de donnée
	up.Id = up.Save()
	// retourne l'objet en JSON
	renderJson(w, up)
}
