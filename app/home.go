package app

import (
	"log"
)

type PageHome struct {
	Images []UserImage
	PageWeb
}

func (ph PageHome) View() Page {

	log.Println("Home appelé")

	// surcharge de la variable d'affichage
	Templ = "index"

	ph.Title = "Coucouc de la forêt"

	ph.MainClass = "accueil"

	return ph
}
