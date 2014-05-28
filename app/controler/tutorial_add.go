package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
	"net/http"
	. "strconv"
)

// fonction public
// permet d'afficher le formulaire de création d'un tutoriel
func (pt *PageTutorial) ViewAdd() {

	log.Println("ViewAdd appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial_add"

	pt.Title = "Créer un nouveau tutoriel"
	pt.MainClass = "forum_add" // même type d'affichage que le formulaire forum
	pt.Categories = t.GetAllCategories(0)
	pt.RenderHtml = true
}

// fonction public
// permet de valider les éléments du formulaire
func (pt PageTutorial) ValidateForm(r *http.Request) (t Tutorial, v bool) {

	log.Println("ValidateForm appelé")

	t.Title = r.PostFormValue("post-nom")
	t.Text = r.PostFormValue("post-contenu")
	t.TutorialCategoryId, _ = ParseInt(r.PostFormValue("post-cat"), 0, 64)
	t.IsOnline = 1 // permet de mettre en ligne automatiquement la quesion

	// on initialise v à true
	v = true
	// on vérifi les champs qui sont présents
	if r.PostFormValue("post-nom") == "" {
		v = false
	}

	if r.PostFormValue("post-contenu") == "" {
		v = false
	}

	return
}
