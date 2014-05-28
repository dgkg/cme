package handler

import (
	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/model"
	. "strconv"
)

// Public function
// permet de récupérer une liste de projets
func (h *Handlers) UserProjectAjaxHandler() (m string) {
	// récupération de la variable id
	vars := mux.Vars(h.R)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	//page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 {
		// envoie un message d'erreur
		return "error"
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCat(3)
		// encodage en json des projets
		b, err := json.Marshal(ups)
		if err != nil {
			// envoie un message d'erreur
			return "error"
		}
		return string(b)
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
func (h *Handlers) UserProjectMoreAjaxHandler() (m string) {
	// récupération de la variable id
	vars := mux.Vars(h.R)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if page == 0 {
		// envoie un message d'erreur
		return "error"
	} else {

		var up UserProject
		ups := up.getListPaged(page, 3)
		// encodage en json des projets
		b, err := json.Marshal(ups)
		if err != nil {
			// envoie un message d'erreur
			return "error"
		}
		return string(b)
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
// et avec une catégorie sélectionnée
func (h *Handlers) UserProjectMoreInCatAjaxHandler() (m string) {
	// récupération de la variable id
	vars := mux.Vars(h.R)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 || page == 0 {
		// envoie un message d'erreur
		return "error"
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCatPaged(page, 3)
		// encodage en json des projets
		b, err := json.Marshal(ups)
		if err != nil {
			// envoie un message d'erreur
			return "error"
		}
		return string(b)
	}
}

// permet de créer un userProject à partir d'un formulaire HTML
// et permet d'uploader une image associée
// et permet de croper l'image en fonction de sa vignette
func (h *Handlers) UPCreateAjaxHandler() (m string) {

	// création d'un user project
	var up UserProject
	// récupère l'id du user
	UserId, err := ParseInt(h.R.PostFormValue("id_user"), 0, 64)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	up.UserId = UserId
	// crécuprère l'id de la cat
	up.UserProjectCategoryId, err = ParseInt(h.R.PostFormValue("id_cat"), 0, 64)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	// récupère le titre
	up.Title = sanitize.HTML(h.R.PostFormValue("title"))
	// récupère la description
	up.Description = sanitize.HTML(h.R.PostFormValue("Description"))
	// permet d'uploader l'image
	up.Url, err = uploadImage(URL_PROJECT_IMAGES, h.R)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	// permet de cropper l'image qui viens d'être uploadée
	err = cropImage(up.Url, 300, 300)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	// finit de créer l'objet
	// définit le projet non visible
	up.IsOnline = 0
	// sauvegarde dans la base de donnée
	up.Id = up.Save()
	// retourne l'objet en JSON
	b, err := json.Marshal(up)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	return string(b)
}
