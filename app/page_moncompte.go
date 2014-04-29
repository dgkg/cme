package app

import (
	"log"
)

type PageCompte struct {
	PageWeb
}

func (pc PageCompte) View() Page {

	log.Println("Mon Compte appelé")

	Templ = "user_compte"

	pc.Title = "Votre compte"
	pc.MainClass = "user_compte"
	pc.RenderHtml = true

	return pc
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageCompte) IsHtmlRender() bool {
	return p.RenderHtml
}
