package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
)

type PageHome struct {
	Categories []UserProjectCategory
	Images     []UserProjectImage
	News       []News
	Projets    []UserProject
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageHome() *PageHome {
	return new(PageHome)
}

func (ph *PageHome) View() {
	log.Println("Accueil appelé")
	Templ = "index"
	ph.Title = "Accueil | Communauté Maryse-Éloy"
	ph.MainClass = "accueil"
	ph.RenderHtml = true
	ph.getProjects()
	ph.getCategories()
	ph.getNews(2)
}

// permet de récupérer les news
func (ph *PageHome) getNews(numb int) {
	var n News
	ph.News = n.GetList(numb)
}

// permet de récupérer les projets
func (ph *PageHome) getProjects() {
	var up UserProject
	ph.Projets = up.GetAllProjects()
}

// permet de récupérer les catégories qui ont des projets
func (ph *PageHome) getCategories() {
	log.Println("getCategories appelé")
	var uic UserProjectCategory
	ph.Categories = uic.GetCategoriesProjects()
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageHome) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageHome) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
	return
}
