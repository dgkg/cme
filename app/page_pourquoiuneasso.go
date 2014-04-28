package app

import (
	"log"
)

type PagePourquoiUneAsso struct {
	PageWeb
}

func (ppua PagePourquoiUneAsso) View() Page {

	log.Println("Pourquoi une association? appelé")

	Templ = "pourquoi_une_association"

	ppua.Title = "Pourquoi une association?"
	ppua.MainClass = "pageinfo"
	ppua.RenderHtml = true

	return ppua
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PagePourquoiUneAsso) IsHtmlRender() bool {
	return p.RenderHtml
}
