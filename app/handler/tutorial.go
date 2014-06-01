package handler

import (
	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/model"
	"log"
	"net/http"
	. "strconv"
	"time"
)

func TutoEditHandler(w http.ResponseWriter, r *http.Request) {

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	id := vars["id"]

	pt := new(PageTutorial)
	pt.ViewEdit(id)

	//insersion dans l'interface Page
	var p Page
	p = pt
	Render(w, p, r)
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
	render(w, p, r)
}

// affichage de la liste des tutos
func TutoHandler(w http.ResponseWriter, r *http.Request) {

	pt := new(PageTutorialList)
	value := r.FormValue("p")

	if value == "" {
		pt.View()
	} else {
		pt.ViewPaged(value)
	}

	//insersion dans l'interface Page
	var p Page
	p = pt
	Render(w, p, r)
}

// affichage d'une catégorie d'un tutoriel
func TutoCatHandler(w http.ResponseWriter, r *http.Request) {
	pt := new(PageTutorialList)

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	value := r.FormValue("p")

	if value == "" {
		pt.ViewCategory(category)
	} else {
		pt.ViewCategoryPaged(category, value)
	}

	//insersion dans l'interface Page
	var p Page
	p = pt
	Render(w, p, r)
}

// affichage de la recherche de tutos
func TutoSearchHandler(w http.ResponseWriter, r *http.Request) {
	pt := new(PageTutorialList)

	q := r.FormValue("q")
	if q == "" {
		pt.View()
	} else {
		pt.ViewSearch(q)
	}

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
		return "error"
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
		u = u.GetById()
		// permet de convertir la date de la personne qui a posté la réponse
		t := time.Now()
		date := t.Format(dateLayout)
		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := u.FirstName + " " + u.LastName + ":::" + date + ":::" + tp.Text + ":::" + Itoa(int(tp.Id))
		return commData
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
	return commData
}

// affichage du formulaire d'ajout d'un tuto
func TutoAddHandler(w http.ResponseWriter, r *http.Request) {

	pt := new(PageTutorial)

	t, v := pt.ValidateForm(h.R)

	if v {
		log.Print("VALIDE!!")
		t.Save()
	} else {
		log.Print("NON VALIDE!!")
	}

	pt.ViewAdd()

	//insersion dans l'interface Page
	var p Page
	p = pt
	render(w, p, r)
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func SubmitTutorialHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("titre_post") == "" ||
		r.PostFormValue("categorie_post") == "" ||
		r.PostFormValue("contenu_post") == "" ||
		r.PostFormValue("user_id") == "" {
		return "error"
	} else {

		var t Tutorial
		t.Title = r.PostFormValue("titre_post")
		t.TutorialCategoryId, _ = ParseInt(r.PostFormValue("categorie_post"), 0, 64)
		t.Text = r.PostFormValue("contenu_post")
		t.UserId, _ = ParseInt(r.PostFormValue("user_id"), 0, 64)
		t.Id, _ = ParseInt(r.PostFormValue("tutorial_id"), 0, 64)
		t.IsOnline = 1
		t.Id = t.Save()
		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := Itoa(int(t.Id)) + ":::" + "All good"
		return commData
	}
}
