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

	pu := new(PageUserList)
	pu.View()

	//insersion dans l'interface Page
	var p Page
	p = pu
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
