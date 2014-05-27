package app

import (
	"github.com/gorilla/mux"
	. "github.com/konginteractive/cme/app/model"
	"log"
	"net/http"
	. "strconv"
)

func TutoEditHandler(w http.ResponseWriter, r *http.Request) {

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	id := vars["id"]

	pt := new(PageTutorial)
	pt.ViewEdit(id)

	//insersion dans l'interface Page
	var p Page
	p = pt
	Render(w, p, r)
}

// fonction public
// permet d'afficher une question avec la liste des réponses
func (pt *PageTutorial) ViewEdit(id string) {
	log.Println("ViewEdit appelé")
	// surcharge de la variable d'affichage
	Templ = "tutorial_add"
	var t Tutorial
	pt.MainClass = "forum_add"
	pt.Tutorial.Id, _ = ParseInt(id, 0, 64)
	pt.Tutorial = pt.Tutorial.getById()
	pt.Categories = t.getAllCategories(0)
	pt.Title = "Éditer : " + pt.Tutorial.Title
	pt.RenderHtml = true
}
