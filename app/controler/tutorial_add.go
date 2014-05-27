package controler

import (
	//"github.com/gorilla/mux"
	"log"
	//"math"
	"fmt"
	. "github.com/konginteractive/cme/app/model"
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
	pt.MainClass = "forum_add" // même type d'affichage que le formulaire forum
	pt.Categories = t.getAllCategories(0)
	pt.RenderHtml = true
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func SubmitTutorialHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("titre_post") == "" ||
		r.PostFormValue("categorie_post") == "" ||
		r.PostFormValue("contenu_post") == "" ||
		r.PostFormValue("user_id") == "" {
		fmt.Fprint(w, "error")
	} else {

		var t Tutorial
		t.Title = r.PostFormValue("titre_post")
		t.TutorialCategoryId, _ = ParseInt(r.PostFormValue("categorie_post"), 0, 64)
		t.Text = r.PostFormValue("contenu_post")
		t.UserId, _ = ParseInt(r.PostFormValue("user_id"), 0, 64)
		t.Id, _ = ParseInt(r.PostFormValue("tutorial_id"), 0, 64)
		t.IsOnline = 1
		t.Id = t.Save()
		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := Itoa(int(t.Id)) + ":::" + "All good"
		fmt.Fprint(w, commData)
	}
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
