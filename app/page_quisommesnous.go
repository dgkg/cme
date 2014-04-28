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
	ppua.RenderHtml = false

	return pqsn
}
