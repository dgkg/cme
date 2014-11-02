package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
	. "strconv"
)

type PageCompte struct {
	PageWeb
	User
	Graduations []string
}

func (pc *PageCompte) View(id int64) {

	log.Println("Mon Compte appelé")

	Templ = "user_compte"

	pc.Title = "Votre compte"
	pc.MainClass = "user_compte"

	var u User
	u.Id = id
	u = u.GetById()

	pc.Text = u.Text
	pc.Facebook = u.Facebook
	pc.Twitter = u.Twitter
	pc.LinkedIn = u.LinkedIn
	pc.Graduation = u.Graduation
	pc.FirstName = u.FirstName
	pc.LastName = u.LastName
	pc.Id = u.Id
	pc.PhotoCover = u.PhotoCover

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
