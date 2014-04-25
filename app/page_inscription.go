package app

import (
	"log"
)

type PageInscription struct {
	PageWeb
}

func (pi PageInscription) View() Page {

	log.Println("Inscription appelé")

	Templ = "inscription"

	pi.Title = "Inscription"
	pi.MainClass = "inscription"

	return pi
}
