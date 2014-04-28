package app

import (
	"log"
)

type PageQuiSommesNous struct {
	PageWeb
}

func (pqsn PageQuiSommesNous) View() Page {

	log.Println("Qui sommes-nous? appelé")

	Templ = "qui_sommes_nous"

	pqsn.Title = "Qui sommes-nous?"
	pqsn.MainClass = "pageinfo"
	pqsn.RenderHtml = true

	return pqsn
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageQuiSommesNous) IsHtmlRender() bool {
	return p.RenderHtml
}
