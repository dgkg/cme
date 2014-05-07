package app

import (
	"log"
	"net/http"
)

type PagePourquoiUneAsso struct {
	PageWeb
}

// affichage de la page pourquoi une association
func PourquoiUneAssoHandler(w http.ResponseWriter, r *http.Request) {

	ppua := new(PagePourquoiUneAsso)
	ppua.View()

	//insersion dans l'interface Page
	var p Page
	p = ppua
	Render(w, p, r)
}

// fonction pour permettre de créer une page
func CreatePagePourquoiUneAsso() *PagePourquoiUneAsso {
	return new(PagePourquoiUneAsso)
}

func (ppua *PagePourquoiUneAsso) View() {

	log.Println("Pourquoi une association? appelé")

	Templ = "pourquoi_une_association"

	ppua.Title = "Pourquoi une association?"
	ppua.MainClass = "pageinfo"
	ppua.RenderHtml = true

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PagePourquoiUneAsso) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PagePourquoiUneAsso) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
