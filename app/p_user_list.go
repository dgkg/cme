package app

import (
	//"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PageUserList struct {
	Users     []User
	PagesList []Paginate
	PageWeb
}

// affichage de la liste des étudants
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	pul := new(PageUserList)
	pul.View()
	pul.render(w, r)
}

// affichage de la recherche dans la liste des étudants
func StudentSearchHandler(w http.ResponseWriter, r *http.Request) {
	pul := new(PageUserList)
	q := r.FormValue("q")
	if q == "" {
		pul.View()
	} else {
		pul.ViewSearch(q)
	}
	pul.render(w, r)
}

func (pul *PageUserList) render(w http.ResponseWriter, r *http.Request) {
	//insersion dans l'interface Page
	var p Page
	p = pul
	Render(w, p, r)
}

func (pu *PageUserList) View() {

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

}

// fonction public
// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les noms
// et va chercher dans les prénoms
func (pul *PageUserList) ViewSearch(q string) {

	log.Println("ForumViewSearch appelé")

	// surcharge de la variable d'affichage
	Templ = "user"

	pul.Title = "Users"
	pul.MainClass = "eleves"
	pul.SearchText = q
	pul.getSearch()

	//pul.Categories = u.getAllCategories(0)
	pul.RenderHtml = true
	//pul.injectDataToDisplay(pul.Users)

}

// getSearch permet d'insérer une liste d'utilisateurs
// en fonction de la recherche q
func (pul *PageUserList) getSearch() {
	var u User
	pul.Users = u.search(pul.SearchText)
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageUserList) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageUserList) SetSessionData(u User) (v bool) {
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
