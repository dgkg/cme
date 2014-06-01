package handler

import (
	. "github.com/konginteractive/cme/app/controler"
)

// affichage de la home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ph := new(PageHome)
	ph.View()
	render(w, p, r)
}

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	nbProjetsACharger := r.PostFormValue("nbProjets")
	return nbProjetsACharger
}
