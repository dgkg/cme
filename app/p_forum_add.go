package app

import (
	//"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	. "strconv"
)

// affichage du formulaire du forum
func ForumAddHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForum)
	f, v := pf.ValidateForm(r)
	if v {
		log.Print("VALIDE!!")
		f.Save()
	} else {
		log.Print("NON VALIDE!!")
	}
	//insersion dans l'interface Page
	var p Page
	p = pf
	pf.ViewAdd()
	Render(w, p, r)
}

// fonction public
// permet d'afficher le formulaire de création d'une question du formulaire
func (pf *PageForum) ViewAdd() {

	log.Println("ViewAdd appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum_add"

	pf.Title = "Créer un nouveau sujet"
	pf.MainClass = "forum_add"
	pf.Forum = make([]Forum, 1)
	pf.Categories = f.getAllCategories()
	pf.RenderHtml = true

}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("titre_post") == "" ||
		r.PostFormValue("resolu_post") == "" ||
		r.PostFormValue("categorie_post") == "" ||
		r.PostFormValue("contenu_post") == "" {

		fmt.Fprint(w, "error")
	} else {

		var f Forum

		isSolved, _ := ParseInt(r.PostFormValue("resolu_post"), 0, 64)
		idCat, _ := ParseInt(r.PostFormValue("categorie_post"), 0, 64)

		f.Title = r.PostFormValue("titre_post")
		f.ForumCategoryId = idCat
		f.IsSolved = isSolved
		f.Text = r.PostFormValue("contenu_post")
		f.IsOnline = 1
		f.Save()
	}
}

// fonction public
// permet de valider les éléments du formulaire
func (pf PageForum) ValidateForm(r *http.Request) (f Forum, v bool) {

	log.Println("ValidateForm appelé")

	f.Title = r.PostFormValue("post-nom")
	f.Text = r.PostFormValue("post-contenu")
	f.ForumCategoryId, _ = ParseInt(r.PostFormValue("post-cat"), 0, 64)
	f.IsOnline = 1 // permet de mettre en ligne automatiquement la quesion

	return
}
