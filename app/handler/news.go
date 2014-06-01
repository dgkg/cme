package handler

import (
	. "github.com/konginteractive/cme/app/controler"
	"net/http"
)

// affichage d'une news unique
func NewsUniqueHandler(w http.ResponseWriter, r *http.Request) {
	pn := new(PageNews)
	pn.View()
	//insersion dans l'interface Page
	var p Page
	p = pn
	render(w, p, r)
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
	render(w, p, r)
}
