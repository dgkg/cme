package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "strconv"
)

func ForumEditHandler(w http.ResponseWriter, r *http.Request) {

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	id := vars["id"]

	pf := new(PageForum)
	pf.ViewEdit(id)

	//insersion dans l'interface Page
	var p Page
	p = pf
	Render(w, p, r)
}

// fonction public
// permet d'afficher une question avec la liste des réponses
func (pfp *PageForum) ViewEdit(id string) {

	log.Println("ViewEdit appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum_add"

	pfp.MainClass = "forum_add"
	pfp.Forum.Id, _ = ParseInt(id, 0, 64)
	pfp.Forum = pfp.Forum.getById()
	pfp.Categories = f.getAllCategories()

	pfp.Title = "Éditer : " + pfp.Forum.Title

	pfp.RenderHtml = true
}
