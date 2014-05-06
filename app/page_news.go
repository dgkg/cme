package app

import (
	"log"
	"math"
	. "strconv"
)

// Helper de vue
type NewsViewHelper struct {
	CategoryTitle string
	News
}

type PageNews struct {
	Categories []NewsCategory
	PagesList  []Paginate
	News       []News
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

	pn.Title = "Actualités de CME"
	pn.MainClass = "news"

	var n News

	pn.News = n.getList()
	pn.PagesList = pn.createPaginate()
	pn.injectDataToDisplay()
	pn.RenderHtml = false

}

// fonction public
// permet d'afficher la liste des questions du forum
// avec la fonction de pagination
func (pn *PageNews) ViewPaged(page string) {

	log.Println("ViewPaged appelé : " + page)

	var n News

	// surcharge de la variable d'affichage
	Templ = "news"

	pagePosition, _ := ParseInt(page, 0, 64)

	pn.Title = "Actualités de CME"
	pn.MainClass = "news"
	pn.News = n.getListPaged(pagePosition)
	pn.Categories = n.getAllCategories()
	pn.PagesList = pn.createPaginate()
	pn.RenderHtml = true
	pn.injectDataToDisplay()

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

// fonction privée
// fonction pour créer la pagination
func (pn PageNews) createPaginate() []Paginate {
	var n News
	elTotal := n.count()

	nb := elTotal / maxElementsInPage
	mf := int(math.Floor(float64(nb)))
	p := make([]Paginate, nb)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
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
