package handler

import (
	"encoding/json"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/helper"
	. "github.com/konginteractive/cme/app/model"
	. "strconv"
)

// UPICreateAjaxHandler permet de créer un User Project Image depuis un formulaire en AJAX
// paramètres de UserProjectImage
// UserProjectImage.ProjectId : id_user_project
// UserProjectImage.Title : title
// UserProjectImage.Description : description
// UserProjectImage.Url : URL_PROJECT_IMAGES + image
// par défaut l'image est visible par tous
func UPICreateAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// création d'une image dans user project
	var upi UserProjectImage
	// récupération des variables
	upi.ProjectId, _ = ParseInt(r.PostFormValue("id_user_project"), 0, 64)
	upi.Title = r.PostFormValue("title")
	upi.Description = r.PostFormValue("description")
	// permet d'uploader l'image
	imgPath, err := UploadImage(URL_PROJECT_IMAGES, h.R)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	// permet de cropper l'image qui viens d'être uploadée
	err = CropImage(imgPath, 300, 300)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	// finit de créer l'objet
	upi.Url = imgPath
	// l'image est par défaut visible
	upi.IsOnline = 1
	// sauvegarde dans la base de donnée
	upi.Id = upi.Save()
	// retourne l'objet en json
	// retourne l'objet en JSON
	b, err := json.Marshal(upi)
	if err != nil {
		// envoie un message d'erreur
		return "error"
	}
	return string(b)
}
