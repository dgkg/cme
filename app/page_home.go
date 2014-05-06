package app

import (
	"log"
)

type PageHome struct {
	Images []UserImage
	News []News
	Projets []UserImage
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageHome() *PageHome {
	return new(PageHome)
}

func (ph *PageHome) View() {

	log.Println("Accueil appelé")

	// surcharge de la variable d'affichage
	Templ = "index"

	ph.Title = "Accueil | Communauté Maryse-Éloy"

	var uimg UserImage

	ph.Projets = uimg.getProjets()

	ph.MainClass = "accueil"
	ph.RenderHtml = true

	var n News
	ph.News = n.getList(2)
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageHome) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageHome) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
