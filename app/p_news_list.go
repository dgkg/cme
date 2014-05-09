package app

import (
	"log"
	"math"
	"net/http"
	. "strconv"
)

type PageNewsList struct {
	Categories []NewsCategory
	PagesList  []Paginate
	News       []News
	PageWeb
}

// affichage de la liste des news
func NewsHandler(w http.ResponseWriter, r *http.Request) {

	pn := new(PageNewsList)

	value := r.FormValue("p")
	if value == "" {
		pn.View()
	} else {
		pn.ViewPaged(value)
	}

	//insersion dans l'interface Page
	var p Page
	p = pn
	Render(w, p, r)
}

func (pn *PageNewsList) View() {

	log.Println("Actualités appelé")

	// surcharge de la variable d'affichage
	Templ = "news"

	pn.Title = "Actualités de CME"
	pn.MainClass = "news"

	var n News

	pn.News = n.getList(3)
	pn.PagesList = pn.createPaginate()
	pn.injectDataToDisplay()
	pn.RenderHtml = false

}

// fonction public
// permet d'afficher la liste des questions du forum
// avec la fonction de pagination
func (pn *PageNewsList) ViewPaged(page string) {

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

func (pn *PageNewsList) injectDataToDisplay() {

	// permet de récupérer le nom prénom de la personne qui a posté la question
	var u User

	// permet de convertir la date de la personne qui a posté la question
	for key, _ := range pn.News {
		t := pn.News[key].CreatedAt
		pn.News[key].CreatedAtString = t.Format(dateLayout)

		u.Id = pn.News[key].UserId
		u = u.getById()
		pn.News[key].UserName = u.FirstName + " " + u.LastName
		//log.Println(u.FirstName)
	}
}

// fonction privée
// fonction pour créer la pagination
func (pn PageNewsList) createPaginate() []Paginate {
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
