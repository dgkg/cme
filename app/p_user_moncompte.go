package app

import (
	"fmt"
	"log"
	"net/http"
	."strconv"
)

type PageCompte struct {
	PageWeb
	User
	Graduations []string
}

func MonCompteHandler(w http.ResponseWriter, r *http.Request) {

	pc := new(PageCompte)

	session, err := store.Get(r, "cme_connecte")

	var u User

	if err == nil && session.Values["id"] != nil {
		u.Id = session.Values["id"].(int64)
	}

	pc.View(u.Id)

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
	sectionRecue := r.PostFormValue("section")

	if sectionRecue == "" {

		fmt.Fprint(w, "error")

	} else {

		var u User

		u.Id, _ = ParseInt(r.PostFormValue("idUser"), 0, 64)
		//log.Println(u.Id)

		switch sectionRecue {

			case "savePhoto" :
				//	@todo : Recevoir la photo +
				//	l'héberger dans public/img/ +
				//	retourner une confirmation

			case "saveBio" :
				u.Text = r.PostFormValue("biographie")
				u.SaveBio()

			case "saveSocial" :
				u.Facebook = r.PostFormValue("facebook")
				u.Twitter = r.PostFormValue("twitter")
				u.LinkedIn = r.PostFormValue("linkedin")
				u.SaveSocial()

			case "saveGraduation" :
				u.Graduation = r.PostFormValue("graduation")
				u.SaveGrad()

			case "saveNomUtilisateur" :
				u.FirstName = r.PostFormValue("prenom")
				u.LastName = r.PostFormValue("nom")
				u.SaveUserName()

			case "supprimerCompte" :
				u.DeleteAccount()

		}

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

func (pc *PageCompte) View( id int64 ) {

	log.Println("Mon Compte appelé")

	Templ = "user_compte"

	pc.Title = "Votre compte"
	pc.MainClass = "user_compte"

	var u User
	u.Id = id
	u = u.getById()

	pc.Text = u.Text
	pc.Facebook = u.Facebook
	pc.Twitter = u.Twitter
	pc.LinkedIn = u.LinkedIn
	pc.Graduation = u.Graduation
	pc.FirstName = u.FirstName
	pc.LastName = u.LastName

	pc.Graduations = make([]string, 20)

	var anneeGrad int = 2014

	for key, _ := range pc.Graduations {
		pc.Graduations[key] = Itoa(anneeGrad)
		anneeGrad--
	}


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
	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
	return
}
