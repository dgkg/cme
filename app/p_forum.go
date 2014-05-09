package app

import (
	//"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "strconv"
)

type PageForum struct {
	Categories []ForumCategory
	Forum      Forum
	PageWeb
}

// affichage d'une question du forum
func ForumPostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("Id appelé : " + id)

	pfp := new(PageForum)
	pfp.Forum.Id, _ = ParseInt(id, 0, 64)
	pfp.View()

	//insersion dans l'interface Page
	var p Page
	p = pfp
	Render(w, p, r)
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
	u = u.getById()
	pfp.Forum.UserName = u.FirstName + " " + u.LastName

	// permet de convertir la date de la personne qui a posté la question
	t := pfp.Forum.CreatedAt
	pfp.Forum.CreatedAtString = t.Format(dateLayout)

	lenPosts := len(pfp.Forum.Posts)

	for i := 0; i < lenPosts; i++ {
		// permet de récupérer le nom prénom de la personne qui a posté la réponse
		u.Id = pfp.Forum.Posts[i].UserId
		u = u.getById()
		pfp.Forum.Posts[i].UserName = u.FirstName + " " + u.LastName

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
	pfp.Forum = pfp.Forum.getById()
	pfp.Forum.Posts = pfp.Forum.getPost()
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
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
