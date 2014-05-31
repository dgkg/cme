package handler

import (
	"github.com/gorilla/mux"
	. "github.com/konginteractive/cme/app/controler"
)

// affichage d'une question du forum
func (h *Handlers) pageSimpleHandler() (p Page) {
	// récupération de la variable id
	vars := mux.Vars(h.R)
	pageUrl := vars["pageUrl"]
	// surcharge de la variable d'affichage
	Templ = "simple"
	// initialisation de l'objet PageSimple
	var ps PageSimple
	ps.RenderHtml = true
	ps.Url = pageUrl
	ps.MainClass = "pageinfo"
	ps.View()
	//insersion dans l'interface Page
	p = ps
	return p
}
