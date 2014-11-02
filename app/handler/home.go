package handler

import (
	. "github.com/konginteractive/cme/app/controler"
	"net/http"
)

// affichage de la home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ph := new(PageHome)
	ph.View()
	var p Page
	p = ph
	render(w, p, r)
}

// fonction qui sert vraiment à rien !
func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	//nbProjetsACharger := r.PostFormValue("nbProjets")
	//return nbProjetsACharger
}
