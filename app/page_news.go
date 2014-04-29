package app

import (
	"log"
)

// Helper de vue
type NewsViewHelper struct {
	CategoryTitle string
	News
}

type PageNews struct {
	PagesList []Paginate
	News      []News
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageNews() *PageNews {
	return new(PageNews)
}

func (pn *PageNews) View() {

	log.Println("Actualités appelé")

	// surcharge de la variable d'affichage
	Templ = "news"

	pn.Title = "News"
	pn.MainClass = "news"

	var n News

	pn.News = n.getList()

	pn.injectDataToDisplay()

	// pagination
	pn.PagesList = make([]Paginate, 5)

	pn.PagesList[0].Title = "1"
	pn.PagesList[0].Url = "/forum/page/1"

	pn.PagesList[1].Title = "2"
	pn.PagesList[1].Url = "/forum/page/2"

	pn.PagesList[2].Title = "3"
	pn.PagesList[2].Url = "/forum/page/3"

	pn.PagesList[3].Title = "4"
	pn.PagesList[3].Url = "/forum/page/4"

	pn.PagesList[4].Title = "5"
	pn.PagesList[4].Url = "/forum/page/5"
	pn.RenderHtml = false

}

func (pn *PageNews) injectDataToDisplay() {

	// permet de récupérer le nom prénom de la personne qui a posté la question
	var u User

	// permet de convertir la date de la personne qui a posté la question
	for key, _ := range pn.News {
		t := pn.News[key].CreatedAt
		pn.News[key].CreatedAtString = t.Format(dateLayout)

		u.Id = pn.News[key].UserId
		u = u.getById()
		pn.News[key].UserName = u.FirstName + " " + u.LastName
		log.Println(u.FirstName)
	}
}

func getNews() {

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
