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
