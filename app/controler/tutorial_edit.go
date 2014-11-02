package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
	. "strconv"
)

// fonction public
// permet d'afficher une question avec la liste des réponses
func (pt *PageTutorial) ViewEdit(id string) {
	log.Println("ViewEdit appelé")
	// surcharge de la variable d'affichage
	Templ = "tutorial_add"
	var t Tutorial
	pt.MainClass = "forum_add"
	pt.Tutorial.Id, _ = ParseInt(id, 0, 64)
	pt.Tutorial = pt.Tutorial.GetById()
	pt.Categories = t.GetAllCategories(0)
	pt.Title = "Éditer : " + pt.Tutorial.Title
	pt.RenderHtml = true
}
