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

func (pn PageNews) View() Page {

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

	return pn
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
