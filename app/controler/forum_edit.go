package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
	. "strconv"
)

// fonction public
// permet d'afficher une question avec la liste des réponses
func (pfp *PageForum) ViewEdit(id string) {

	log.Println("ViewEdit appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum_add"

	pfp.MainClass = "forum_add"
	pfp.Forum.Id, _ = ParseInt(id, 0, 64)
	pfp.Forum = pfp.Forum.GetById()
	pfp.Categories = f.GetAllCategories(0)

	pfp.Title = "Éditer : " + pfp.Forum.Title

	pfp.RenderHtml = true
}
