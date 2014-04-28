package app

import (
	"log"
)

type PageConnexion struct {
	PageWeb
}

func (pc PageConnexion) View() Page {

	log.Println("Connexion appelé")

	Templ = "connexion"

	pc.Title = "Connexion"
	pc.MainClass = "connexion"

	pc.RenderHtml = false

	return pc
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageConnexion) IsHtmlRender() bool {
	return p.RenderHtml
}
