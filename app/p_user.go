package app

import (
	//"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PageUser struct {
	User
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageUser() *PageUser {
	return new(PageUser)
}

// affichage de la fiche étudiant
func StudentFicheHandler(w http.ResponseWriter, r *http.Request) {

	pu := new(PageUser)
	pu.View()

	//insersion dans l'interface Page
	var p Page
	p = pu
	Render(w, p, r)
}

func (pu *PageUser) View() {

	log.Println("Fiche appelé")

	Templ = "user_fiche"
	pu.Title = "Fiche de Henri Lepic"
	pu.MainClass = "eleve_fiche"

	pu.RenderHtml = true
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageUser) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageUser) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
