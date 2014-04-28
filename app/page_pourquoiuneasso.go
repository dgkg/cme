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
