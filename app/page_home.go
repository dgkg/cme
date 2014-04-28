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
	ph.RenderHtml = true
	return ph
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageHome) IsHtmlRender() bool {
	return p.RenderHtml
}
