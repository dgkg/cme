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

type PageTutorial struct {
	Categories []TutorialCategory
	Tutorial   Tutorial
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageTutorial() *PageTutorial {
	return new(PageTutorial)
}

// affichage d'un tuto
func TutoPostHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]
	// initialisation de l'objet PageTutorial
	pt := new(PageTutorial)
	pt.Tutorial.Id, _ = ParseInt(id, 0, 64)
	pt.View()
	//insersion dans l'interface Page
	var p Page
	p = pt
	Render(w, p, r)
}

// Public function
// permet d'ajouter un commenaire sur la page tutoriel
func TutorialNouvCommHandler(w http.ResponseWriter, r *http.Request) {
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
		var tp TutorialPost
		tp.TutorialId, _ = ParseInt(r.PostFormValue("val_post_id"), 0, 64)
		tp.UserId, _ = ParseInt(r.PostFormValue("val_auteur_id"), 0, 64)
		tp.Text = sanitize.HTML(r.PostFormValue("val_commentaire"))
		tp.IsOnline = 1
		tp.Id = tp.Save()
		// permet de récuprérer le nom de l'utilisateur
		var u User
		u.Id = tp.UserId
		u = u.getById()
		// permet de convertir la date de la personne qui a posté la réponse
		t := time.Now()
		date := t.Format(dateLayout)
		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := u.FirstName + " " + u.LastName + ":::" + date + ":::" + tp.Text + ":::" + Itoa(int(tp.Id))
		fmt.Fprint(w, commData)
	}
}

// fonction Public
// permet de supprimer un commentaire sur le tuto
func TutorialDelCommHandler(w http.ResponseWriter, r *http.Request) {
	var tp TutorialPost
	tp.Id, _ = ParseInt(r.PostFormValue("id_commentaire"), 0, 64)
	log.Println("TutorialPost " + Itoa(int(tp.Id)) + " supprimé")
	tp.Delete()
	commData := "success"
	fmt.Fprint(w, commData)
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
	pt.Tutorial.UserName = u.getFullName()

	// permet de convertir la date de la personne qui a posté la question
	t := pt.Tutorial.CreatedAt
	pt.Tutorial.CreatedAtString = t.Format(dateLayout)

	lenPosts := len(pt.Tutorial.Posts)
	for i := 0; i < lenPosts; i++ {
		// permet de récupérer le nom prénom de la personne qui a posté la réponse
		u.Id = pt.Tutorial.Posts[i].UserId
		pt.Tutorial.Posts[i].UserName = u.getFullName()
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
	pt.Tutorial = pt.Tutorial.getById()
	pt.Tutorial.Posts = pt.Tutorial.getPost()
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
