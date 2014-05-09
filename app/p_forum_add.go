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
	pf.ViewAdd()

	//insersion dans l'interface Page
	var p Page
	p = pf
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
	pf.Categories = f.getAllCategories(0)
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
