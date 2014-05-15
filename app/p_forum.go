package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	"log"
	"net/http"
	. "strconv"
	"time"
)

type PageForum struct {
	Categories []ForumCategory
	Forum      Forum
	PageWeb
}

// affichage d'une question du forum
func ForumPostHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]
	// initialisation de l'objet PageForum
	pfp := new(PageForum)
	pfp.Forum.Id, _ = ParseInt(id, 0, 64)
	pfp.View()
	//insersion dans l'interface Page
	var p Page
	p = pfp
	Render(w, p, r)
}

// Public function
// permet d'ajouter un commenaire sur une fonction
func ForumNouvCommHandler(w http.ResponseWriter, r *http.Request) {
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if r.PostFormValue("val_commentaire") == "" ||
		r.PostFormValue("val_post_id") == "" ||
		r.PostFormValue("val_auteur_id") == "" ||
		r.PostFormValue("val_auteur_id") == "0" {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {
		// initialise l'objet ForumPost et récupère les données du formulaire
		var fp ForumPost
		fp.ForumId, _ = ParseInt(r.PostFormValue("val_post_id"), 0, 64)
		fp.UserId, _ = ParseInt(r.PostFormValue("val_auteur_id"), 0, 64)
		fp.Text = sanitize.HTML(r.PostFormValue("val_commentaire"))
		fp.IsOnline = 1
		fp.Id = fp.Save()
		// permet de récuprérer le nom de l'utilisateur
		var u User
		u.Id = fp.UserId
		u = u.getById()
		// permet de convertir la date de la personne qui a posté la réponse
		t := time.Now()
		date := t.Format(dateLayout)
		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := u.FirstName + " " + u.LastName + ":::" + date + ":::" + fp.Text + ":::" + Itoa(int(fp.Id))
		fmt.Fprint(w, commData)
	}
}

// fonction Public
// permet de supprimer un commentaire sur une question
func ForumDelCommHandler(w http.ResponseWriter, r *http.Request) {
	var fp ForumPost
	fp.Id, _ = ParseInt(r.PostFormValue("id_commentaire"), 0, 64)
	fp.Delete()
	commData := "success"
	fmt.Fprint(w, commData)
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
