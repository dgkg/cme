package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
)

// fonction public
// permet d'afficher le formulaire de création d'une question du formulaire
func (pf *PageForum) ViewAdd() {

	log.Println("ViewAdd appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum_add"

	pf.Title = "Créer un nouveau sujet"
	pf.MainClass = "forum_add"
	pf.Categories = f.GetAllCategories(0)
	pf.RenderHtml = true

}
