package app

import (
	"fmt"
	"log"
	"net/http"
)

type PageInscription struct {
	PageWeb
}

// affichage du formulaire d'inscription
func InscriptionHandler(w http.ResponseWriter, r *http.Request) {

	pi := new(PageInscription)
	pi.View()

	//insersion dans l'interface Page
	var p Page
	p = pi
	Render(w, p, r)
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func InscFormHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("prenom") == "" ||
		r.PostFormValue("nom") == "" ||
		r.PostFormValue("email") == "" ||
		r.PostFormValue("mdp") == "" {

		fmt.Fprint(w, "error")
	} else {

		var u User

		u.FirstName = r.PostFormValue("prenom")
		u.LastName = r.PostFormValue("nom")
		u.Email = r.PostFormValue("email")
		u.Pass = r.PostFormValue("mdp")
		u.IsOnline = 1
		u.Save()
	}
}

// fonction pour permettre de créer une page
func CreatePageInscription() *PageInscription {
	return new(PageInscription)
}

func (pi *PageInscription) View() {

	log.Println("Inscription appelé")

	Templ = "inscription"

	pi.Title = "Inscription"
	pi.MainClass = "inscription"
	pi.RenderHtml = false
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageInscription) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageInscription) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
	return
}
