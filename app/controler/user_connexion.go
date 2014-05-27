package app

import (
	"github.com/gorilla/sessions"
	. "github.com/konginteractive/cme/app/model"
	"log"
	"net/http"
)

type PageConnexion struct {
	PageWeb
}

// affichage de la connexion
func ConnexionHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	u.Email = r.PostFormValue("login")
	u.Pass = r.PostFormValue("pass")

	var connected bool
	if u.Email != "" && u.Pass != "" {
		connected, _ = u.LoginPassExist()
	}

	pc := new(PageConnexion)
	pc.View()

	if connected {
		session, _ := store.Get(r, "cme_connecte")
		// Set some session values.
		session.Values["id"] = u.Id
		session.Values["name"] = u.FirstName
		session.Options = &sessions.Options{
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		// Save it.
		session.Save(r, w)
		// Set data into current page
		pc.SessIdUser = u.Id
		pc.SessNameUser = u.FirstName
		pc.SessIsLogged = true
		//http.Redirect(w, r, "http://www.google.com", 301)
	}

	//insersion dans l'interface Page
	var p Page
	p = pc
	Render(w, p, r)

}

// DeconnexionHandler permet de supprimer la connexion "cme_connecte"
// en changant la date d'expiration à -1
// puis affiche la home
func DeconnexionHandler(w http.ResponseWriter, r *http.Request) {
	// récupère la cession en cours
	session, _ := store.Get(r, "cme_connecte")
	// modifie sa date d'expiration
	session.Options.MaxAge = -1
	// sauvegarde
	session.Save(r, w)
	// creation de l'objet PageSimple
	var ps PageSimple
	ps.SessIdUser = 0
	ps.SessIsLogged = false
	ps.SessNameUser = ""
	// surcharge du template à utiliser
	Templ = "deconnexion"
	//insersion dans l'interface Page
	var p Page
	p = ps
	Render(w, p, r)
}

func ConnexionPostHandler(w http.ResponseWriter, r *http.Request) {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "cme_connecte")
	// Set some session values.
	session.Values["name"] = "Antoine"
	session.Values["id"] = 1
	session.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	// Save it.
	session.Save(r, w)
}

func ConnexionGet(w http.ResponseWriter, r *http.Request) {

	//session, _ := store.Get(r, "cme_connecte")

	//fmt.Fprintf(w, "Hi there, I love %s!", session.Values["name"])

	// permet d'afficher quelque chose
	//var pc PageConnexion
	//Render(w, pc.View())
}

// fonction pour permettre de créer une page
func CreatePageConnexion() *PageConnexion {
	return new(PageConnexion)
}

func (pc *PageConnexion) View() {

	log.Println("Connexion appelé")

	Templ = "connexion"

	pc.Title = "Connexion"
	pc.MainClass = "connexion"

	pc.RenderHtml = false
}

// fonction permettant de connecter un utilisateur
// retourne un utilisateur
func (p *PageConnexion) Connect(login string, pass string) (u User, err error) {

	u.Email = login
	u.Pass = pass
	//u, err = u.serchUserLoginAndPass()

	if err == nil && u.Id != 0 {

	}
	return
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageConnexion) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageConnexion) SetSessionData(u User) (v bool) {
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
