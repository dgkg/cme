package app

import (
	"log"
)

type PageCompte struct {
	PageWeb
}

func (pc *PageCompte) View() {

	log.Println("Mon Compte appelé")

	Templ = "user_compte"

	pc.Title = "Votre compte"
	pc.MainClass = "user_compte"
	pc.RenderHtml = true
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageCompte) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageCompte) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
