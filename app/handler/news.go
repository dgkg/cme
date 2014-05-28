package handler

import (
	. "github.com/konginteractive/cme/app/controler"
)

// affichage d'une news unique
func (h *Handlers) NewsUniqueHandler() (p Page) {
	pn := new(PageNews)
	pn.View()
	//insersion dans l'interface Page
	p = pn
	return p
}

// affichage de la liste des news
func (h *Handlers) NewsHandler() (p Page) {

	pn := new(PageNewsList)
	value := h.R.FormValue("p")
	if value == "" {
		pn.View()
	} else {
		pn.ViewPaged(value)
	}
	//insersion dans l'interface Page
	p = pn
	return p
}
