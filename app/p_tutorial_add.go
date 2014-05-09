package app

import (
	//"github.com/gorilla/mux"
	"log"
	//"math"
	"net/http"
	. "strconv"
)

// affichage du formulaire d'ajout d'un tuto
func TutoAddHandler(w http.ResponseWriter, r *http.Request) {

	pt := new(PageTutorial)

	t, v := pt.ValidateForm(r)

	if v {
		log.Print("VALIDE!!")
		t.Save()
	} else {
		log.Print("NON VALIDE!!")
	}

	pt.ViewAdd()

	//insersion dans l'interface Page
	var p Page
	p = pt
	Render(w, p, r)
}

// fonction public
// permet d'afficher le formulaire de création d'un tutoriel
func (pt *PageTutorial) ViewAdd() {

	log.Println("ViewAdd appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial_add"

	pt.Title = "Créer un nouveau tutoriel"
	pt.MainClass = "tutorial_add"
	pt.Categories = t.getAllCategories()
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
