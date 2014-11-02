package handler

import (
	"github.com/gorilla/mux"
	. "github.com/konginteractive/cme/app/controler"
	"net/http"
)

// affichage d'une question du forum
func PageSimpleHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
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
	var p Page
	p = ps
	render(w, p, r)
}
