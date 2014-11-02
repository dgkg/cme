package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
)

// PageSimple est un type qui permet d'afficher des pages d'information
type PageSimple struct {
	Simple
	PageWeb `sql:"-"` // Ignore this field
}

// fonction public
// permet d'afficher la page
func (ps *PageSimple) View() {
	//
	log.Println("View appelé")

	var s Simple
	ps.Simple = s.GetData()
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (ps PageSimple) IsHtmlRender() bool {
	return ps.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (ps PageSimple) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		ps.SessIdUser = u.Id
		ps.SessIsLogged = true
		ps.SessNameUser = u.FirstName
		v = true
	} else {
		ps.SessIdUser = 0
		ps.SessIsLogged = false
		ps.SessNameUser = ""
	}
	return
}
