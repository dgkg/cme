package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "strconv"
)

type PageTutorial struct {
	Categories []TutorialCategory
	Tutorial   Tutorial
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageTutorial() *PageTutorial {
	return new(PageTutorial)
}

// affichage d'un tuto
func TutoPostHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]
	// initialisation de l'objet PageTutorial
	tp := new(PageTutorial)
	tp.Tutorial.Id, _ = ParseInt(id, 0, 64)
	tp.View()
	//insersion dans l'interface Page
	var p Page
	p = tp
	Render(w, p, r)
}

// fonction public
// permet d'afficher la liste des tutoriels
func (pts *PageTutorial) View() {
	//
	log.Println("View appelé")
	// surcharge de la variable d'affichage
	Templ = "tutorial_single"
	//
	pts.Title = "Tutoriel"
	pts.MainClass = "tutoriels"
	//pts.Tutorial.Id, _ = ParseInt(idTuto, 0, 64)
	//pts.Tutorial.getData()
	pts.Tutorial.Posts = pts.Tutorial.getPost()
	pts.RenderHtml = true
	//pts.injectDataToDisplay()

	log.Println("le titre du post est : " + pts.Tutorial.Title)

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageTutorial) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageTutorial) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
	return
}
