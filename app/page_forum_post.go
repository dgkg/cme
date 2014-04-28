package app

import (
	"log"
	. "strconv"
)

type PageForumPost struct {
	//PagesList []Paginate
	//News      []NewsViewHelper
	Forum Forum

	PageWeb
}

// fonction public
// permet d'afficher une question avec la liste des réponses
func (pfp PageForumPost) View(id string) Page {

	log.Println("Forum Post appelé")

	log.Println("Id appelé : " + id)

	idForm, _ := ParseInt(id, 0, 64)

	// surcharge de la variable d'affichage
	Templ = "forum_post"

	pfp.Title = "Titre du post! Woohoo!"
	pfp.MainClass = "forum_post"

	pfp.Forum.Id = idForm
	pfp.Forum = pfp.Forum.getById() //pfp.Forum.getPost(idForm)
	pfp.Forum.Posts = pfp.Forum.getPost()
	pfp.RenderHtml = false
	pfp.injectDataToDisplay()

	log.Println("le titre du post est : " + pfp.Forum.Title)

	return pfp
}

// fonction privée
// Permet de retrouver le nom de la personne qui a posté
// Permet aussi de convertir la date de création
// Permet de récupérer le nom de l'utilisateur qui a posté une réponse
// Permet de convertir la date du post de l'utilisateur qui a posté une réponse
func (pfp *PageForumPost) injectDataToDisplay() {

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
