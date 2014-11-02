package controler

import (
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/model"
	"log"
)

type PageUserList struct {
	Users     []User
	PagesList []Paginate
	PageWeb
}

func (pul *PageUserList) View() {

	log.Println("Users appelé")

	// surcharge de la variable d'affichage
	Templ = "user"

	pul.Title = "Users"
	pul.MainClass = "eleves"

	// pagination
	pul.PagesList = make([]Paginate, 5)

	var u User

	pul.Users = u.GetAllUser()
	pul.injectDataToDisplay()

	pul.PagesList[0].Title = "1"
	pul.PagesList[0].Url = "/eleves/page/1"

	pul.PagesList[1].Title = "2"
	pul.PagesList[1].Url = "/eleves/page/2"

	pul.PagesList[2].Title = "3"
	pul.PagesList[2].Url = "/eleves/page/3"

	pul.PagesList[3].Title = "4"
	pul.PagesList[3].Url = "/eleves/page/4"

	pul.PagesList[4].Title = "5"
	pul.PagesList[4].Url = "/eleves/page/5"

	pul.RenderHtml = true

}

// fonction privée
// permet de créer les urls des projets
func (pul *PageUserList) injectDataToDisplay() {
	for key, _ := range pul.Users {

		// Génère l'URL du profil selon son prénom, nom et l'année de graduation
		fn := sanitize.Name(pul.Users[key].FirstName)
		ln := sanitize.Name(pul.Users[key].LastName)
		ug := pul.Users[key].Graduation
		pul.Users[key].Url = "/eleves/" + ug + "/" + fn + "_" + ln

		// Si l'utilisateur n'a pas choisi sa photo de profil,
		// une image générique lui est attribué
		if pul.Users[key].PhotoProfil == "" {
			pul.Users[key].PhotoProfil = "user_inconnu.jpg"
		}
	}
}

// fonction public
// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les noms
// et va chercher dans les prénoms
func (pul *PageUserList) ViewSearch(q string) {

	log.Println("ForumViewSearch appelé")

	// surcharge de la variable d'affichage
	Templ = "user"

	pul.Title = "Users"
	pul.MainClass = "eleves"
	pul.SearchText = q
	pul.getResults()

	//pul.Categories = u.getAllCategories(0)
	pul.RenderHtml = true
	//pul.injectDataToDisplay(pul.Users)

}

// getSearch permet d'insérer une liste d'utilisateurs
// en fonction de la recherche q
func (pul *PageUserList) getResults() {
	var u User
	pul.Users = u.Search(pul.SearchText)
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageUserList) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageUserList) SetSessionData(u User) (v bool) {
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
