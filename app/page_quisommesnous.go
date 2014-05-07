package app

import (
	"log"
	"net/http"
)

type PageQuiSommesNous struct {
	PageWeb
}

// affichage de la page qui sommes nous
func QuiSommesNousHandler(w http.ResponseWriter, r *http.Request) {

	pqsn := new(PageQuiSommesNous)
	pqsn.View()

	//insersion dans l'interface Page
	var p Page
	p = pqsn
	Render(w, p, r)
}

// fonction pour permettre de créer une page
func CreatePageQuiSommesNous() *PageQuiSommesNous {
	return new(PageQuiSommesNous)
}

func (pqsn *PageQuiSommesNous) View() {

	log.Println("Qui sommes-nous? appelé")

	Templ = "qui_sommes_nous"

	pqsn.Title = "Qui sommes-nous?"
	pqsn.MainClass = "pageinfo"
	pqsn.RenderHtml = true

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageQuiSommesNous) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageQuiSommesNous) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
