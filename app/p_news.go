package app

import (
	"log"
	//"math"
	"net/http"
	//. "strconv"
)

// Helper de vue
type NewsViewHelper struct {
	CategoryTitle string
	News
}

type PageNews struct {
	Categories []NewsCategory
	PagesList  []Paginate
	News
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageNews() *PageNews {
	return new(PageNews)
}

// affichage d'une news unique
func NewsUniqueHandler(w http.ResponseWriter, r *http.Request) {

	pn := new(PageNews)

	pn.View()

	//insersion dans l'interface Page
	var p Page
	p = pn
	Render(w, p, r)
}

func (pn *PageNews) View() {

	log.Println("Actualité unique appelé")

	// surcharge de la variable d'affichage
	Templ = "news-unique"

	pn.PageWeb.Title = "Nom de l'article | Actualités de CME"
	pn.MainClass = "news-unique"

	//pn.injectDataToDisplay()
	pn.RenderHtml = false

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageNews) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageNews) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
