package app

import (
	"log"
	"net/http"
	"fmt"
)

type PageCompte struct {
	PageWeb
}

func MonCompteHandler(w http.ResponseWriter, r *http.Request) {

	pc := new(PageCompte)
	pc.View()

	//insersion dans l'interface Page
	var p Page
	p = pc
	Render(w, p, r)

}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func EditCompteHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("section") == "" /* || r.PostFormValue("photo_profil") == "" */ {

		fmt.Fprint(w, "error")
	} else {

		fmt.Fprint(w, "Bonjour")

		// @todo : Recevoir la section + la photo de profil envoyé
		// 		Héberger l'image sur le serveur dans public/img/
		//		 Retourner une confirmation

		/*

		var f Forum

		isSolved, _ := ParseInt(r.PostFormValue("resolu_post"), 0, 64)
		idCat, _ := ParseInt(r.PostFormValue("categorie_post"), 0, 64)

		f.Title = r.PostFormValue("titre_post")
		f.ForumCategoryId = idCat
		f.IsSolved = isSolved
		f.Text = r.PostFormValue("contenu_post")
		f.IsOnline = 1
		f.Save()
		*/
	}
}

func (pc *PageCompte) View() {

	log.Println("Mon Compte appelé")

	Templ = "user_compte"

	pc.Title = "Votre compte"
	pc.MainClass = "user_compte"
	pc.RenderHtml = true
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageCompte) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageCompte) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
