package app

import (
	"log"
)

type PageInscription struct {
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageInscription() *PageInscription {
	return new(PageInscription)
}

func (pi *PageInscription) View() {

	log.Println("Inscription appelé")

	Templ = "inscription"

	pi.Title = "Inscription"
	pi.MainClass = "inscription"
	pi.RenderHtml = false
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageInscription) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageInscription) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
