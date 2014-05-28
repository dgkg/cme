package controler

import (
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/model"
	"log"
)

type PageUser struct {
	User
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageUser() *PageUser {
	return new(PageUser)
}

func (pu *PageUser) View() {

	var up UserProject
	up.UserId = pu.User.Id

	log.Println("Fiche appelé")

	Templ = "user_fiche"
	pu.Title = "Fiche d'" + pu.User.FirstName + " " + pu.User.LastName
	pu.MainClass = "eleve_fiche"
	pu.User.Projects = up.GetAllFromIdUser()
	pu.RenderHtml = true
	pu.injectDataToDisplay()
}

// fonction privée
// permet de créer les urls des projets
func (pu *PageUser) injectDataToDisplay() {
	for key, _ := range pu.Projects {
		pu.Projects[key].Url = "/eleves/" + pu.User.Graduation + "/" + pu.User.FirstName + "_" + pu.User.LastName + "/" + sanitize.Name(pu.Projects[key].Title)
	}

}

// findUser permet de retrouver le user à partir de :
// son année de graduation
// et de son prénom formatté en url
// et de son nom formatté en url
// retourne un booleen permettan de savoir si oui ou non il a été trouvé
func (pu *PageUser) findUser() (isFound bool) {

	listUsers := pu.User.GetUsersByGraduation()
	sfn := pu.User.FirstName
	sln := pu.User.LastName
	for key, _ := range listUsers {
		fn := sanitize.Name(listUsers[key].FirstName)
		ln := sanitize.Name(listUsers[key].LastName)
		//log.Println("sfn : " + sfn + " / fn : " + fn + " / sln : " + sln + " / ln : " + ln)
		if sfn == fn && sln == ln {
			pu.User = listUsers[key]
			isFound = true
		}
	}
	return
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageUser) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageUser) SetSessionData(u User) (v bool) {
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
