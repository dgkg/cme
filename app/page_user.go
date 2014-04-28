package app

import (
	"log"
)

type PageUser struct {
	PagesList []Paginate
	PageWeb
}

func (pu PageUser) View() Page {

	log.Println("Users appelé")

	// surcharge de la variable d'affichage
	Templ = "user"

	pu.Title = "Users"
	pu.MainClass = "eleves"

	// pagination
	pu.PagesList = make([]Paginate, 5)

	pu.PagesList[0].Title = "1"
	pu.PagesList[0].Url = "/eleves/page/1"

	pu.PagesList[1].Title = "2"
	pu.PagesList[1].Url = "/eleves/page/2"

	pu.PagesList[2].Title = "3"
	pu.PagesList[2].Url = "/eleves/page/3"

	pu.PagesList[3].Title = "4"
	pu.PagesList[3].Url = "/eleves/page/4"

	pu.PagesList[4].Title = "5"
	pu.PagesList[4].Url = "/eleves/page/5"

	pu.RenderHtml = true

	return pu
}

func (pu PageUser) ViewFiche() Page {

	log.Println("Fiche appelé")

	Templ = "user_fiche"
	pu.Title = "Fiche de Henri Lepic"
	pu.MainClass = "eleve_fiche"

	pu.RenderHtml = true

	return pu
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageUser) IsHtmlRender() bool {
	return p.RenderHtml
}
