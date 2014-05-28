package controler

import (
	. "github.com/konginteractive/cme/app/model"
	"log"
)

type PageForum struct {
	Categories []ForumCategory
	Forum      Forum
	PageWeb
}

// fonction privée
// Permet de retrouver le nom de la personne qui a posté
// Permet aussi de convertir la date de création
// Permet de récupérer le nom de l'utilisateur qui a posté une réponse
// Permet de convertir la date du post de l'utilisateur qui a posté une réponse
func (pfp *PageForum) injectDataToDisplay() {

	// permet de récupérer le nom prénom de la personne qui a posté la question
	var u User
	u.Id = pfp.Forum.UserId
	pfp.Forum.UserName = u.GetFullName()

	// permet de convertir la date de la personne qui a posté la question
	t := pfp.Forum.CreatedAt
	pfp.Forum.CreatedAtString = t.Format(dateLayout)

	lenPosts := len(pfp.Forum.Posts)
	for i := 0; i < lenPosts; i++ {
		// permet de récupérer le nom prénom de la personne qui a posté la réponse
		u.Id = pfp.Forum.Posts[i].UserId
		pfp.Forum.Posts[i].UserName = u.GetFullName()
		// permet de convertir la date de la personne qui a posté la question
		t = pfp.Forum.Posts[i].CreatedAt
		pfp.Forum.Posts[i].CreatedAtString = t.Format(dateLayout)
	}
}

// fonction public
// permet d'afficher une question avec la liste des réponses
func (pfp *PageForum) View() {

	log.Println("Forum Post appelé")

	// surcharge de la variable d'affichage
	Templ = "forum_post"

	pfp.Title = "Titre du post! Woohoo!"
	pfp.MainClass = "forum_post"
	pfp.Forum = pfp.Forum.GetById()
	log.Println("*--* pfp.Forum *--*")
	log.Println(pfp.Forum)
	log.Println("*--* fin pfp.Forum *--*")
	pfp.Forum.Posts = pfp.Forum.GetPost()
	pfp.RenderHtml = true
	pfp.injectDataToDisplay()

	log.Println("le titre du post est : " + pfp.Forum.Title)

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageForum) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageForum) SetSessionData(u User) (v bool) {
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

/*






*/
/*
type PageTestPost struct {
	Posts []string
	PageWeb
}

// affichage des posts
func TestPostHandler(w http.ResponseWriter, r *http.Request) {

	var pt PageTestPost
	pt.Posts = make([]string, 3)
	pt.Posts[0] = "All good !!!"
	pt.Posts[1] = r.PostFormValue("login")
	pt.Posts[2] = r.PostFormValue("pass")
	fmt.Fprint(w, pt.Posts)

}
*/
