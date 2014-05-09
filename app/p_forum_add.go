package app

import (
	//"github.com/gorilla/mux"
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
	pf.Forums = make([]Forum, 1)
	pf.Categories = f.getAllCategories()
	pf.RenderHtml = true

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
