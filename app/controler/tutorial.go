package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
)

type PageTutorial struct {
	Categories []TutorialCategory
	Tutorial   Tutorial
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageTutorial() *PageTutorial {
	return new(PageTutorial)
}

// fonction privée
// Permet de retrouver le nom de la personne qui a posté
// Permet aussi de convertir la date de création
// Permet de récupérer le nom de l'utilisateur qui a posté une réponse
// Permet de convertir la date du post de l'utilisateur qui a posté une réponse
func (pt *PageTutorial) injectDataToDisplay() {

	// permet de récupérer le nom prénom de la personne qui a posté la question
	var u User
	u.Id = pt.Tutorial.UserId
	pt.Tutorial.UserName = u.GetFullName()

	// permet de convertir la date de la personne qui a posté la question
	t := pt.Tutorial.CreatedAt
	pt.Tutorial.CreatedAtString = t.Format(dateLayout)

	lenPosts := len(pt.Tutorial.Posts)
	for i := 0; i < lenPosts; i++ {
		// permet de récupérer le nom prénom de la personne qui a posté la réponse
		u.Id = pt.Tutorial.Posts[i].UserId
		pt.Tutorial.Posts[i].UserName = u.GetFullName()
		// permet de convertir la date de la personne qui a posté la question
		t = pt.Tutorial.Posts[i].CreatedAt
		pt.Tutorial.Posts[i].CreatedAtString = t.Format(dateLayout)
	}
}

// fonction public
// permet d'afficher la liste des tutoriels
func (pt *PageTutorial) View() {
	//
	log.Println("View appelé")
	// surcharge de la variable d'affichage
	Templ = "tutorial_single"
	//
	pt.Title = "Tutoriel"
	//pt.MainClass = "tutoriels"
	pt.MainClass = "forum_post"
	pt.Tutorial = pt.Tutorial.GetById()
	pt.Tutorial.Posts = pt.Tutorial.GetPost()
	pt.RenderHtml = true
	pt.injectDataToDisplay()

	log.Println("le titre du post est : " + pt.Tutorial.Title)

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageTutorial) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageTutorial) SetSessionData(u User) (v bool) {
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
