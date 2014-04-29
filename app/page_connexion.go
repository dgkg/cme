package app

import (
	"log"
)

type PageConnexion struct {
	PageWeb
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
	u, err = u.serchUserLoginAndPass()

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
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
