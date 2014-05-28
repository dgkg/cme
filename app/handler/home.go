package handler

import (
	. "github.com/konginteractive/cme/app/controler"
)

// affichage de la home
func (h *Handlers) HomeHandler() (p Page) {
	ph := new(PageHome)
	ph.View()
	return p
}

func (h *Handlers) GetProjectsHandler() (m string) {
	nbProjetsACharger := h.R.PostFormValue("nbProjets")
	return nbProjetsACharger
}
