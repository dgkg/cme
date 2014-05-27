package controler

import (
	"fmt"
	. "github.com/konginteractive/cme/app/model"
	"net/http"
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
	imgPath, err := uploadImage(URL_PROJECT_IMAGES, r)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	}
	// permet de cropper l'image qui viens d'être uploadée
	err = cropImage(imgPath, 300, 300)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	}
	// finit de créer l'objet
	upi.Url = imgPath
	// l'image est par défaut visible
	upi.IsOnline = 1
	// sauvegarde dans la base de donnée
	upi.Id = upi.Save()
	// retourne l'objet en json
	RenderJson(w, upi)
}
