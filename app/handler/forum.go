package handler

import (
	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/model"
	"net/http"
	. "strconv"
	"time"
)

// affichage d'une question du forum
func ForumPostHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	id := vars["id"]
	// initialisation de l'objet PageForum
	pf := new(PageForum)
	pf.Forum.Id, _ = ParseInt(id, 0, 64)
	pf.View()
	//insersion dans l'interface Page
	var p Page
	p = pf
	render(w, p, r)
}

// affichage du formulaire du forum
func ForumAddHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForum)
	pf.ViewAdd()
	//insersion dans l'interface Page
	var p Page
	p = pf
	render(w, p, r)
}

func ForumEditHandler(w http.ResponseWriter, r *http.Request) {

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	id := vars["id"]

	pf := new(PageForum)
	pf.ViewEdit(id)

	//insersion dans l'interface Page
	var p Page
	p = pf
	render(w, p, r)
}

// affichage du forum
func ForumHandler(w http.ResponseWriter, r *http.Request) {

	pf := new(PageForumList)

	value := h.R.FormValue("p")
	if value == "" {
		pf.View()
	} else {
		pf.ViewPaged(value)
	}

	//insersion dans l'interface Page
	var p Page
	p = pf
	render(w, p, r)
}

// affichage de la recherche dans le forum
func ForumSearchHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForumList)
	q := h.R.FormValue("q")
	if q == "" {
		pf.View()
	} else {
		pf.ViewSearch(q)
	}

	//insersion dans l'interface Page
	var p Page
	p = pf
	render(w, p, r)
}

// affichage d'une catégorie dans le forum
func ForumCatHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForumList)

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	value := h.R.FormValue("p")

	if value == "" {
		pf.ViewCategory(category)
	} else {
		pf.ViewCategoryPaged(category, value)
	}
	//insersion dans l'interface Page
	var p Page
	p = pf
	render(w, p, r)
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
		return "error"
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
		u = u.GetById()
		// permet de convertir la date de la personne qui a posté la réponse
		t := time.Now()
		date := t.Format(dateLayout)
		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := u.FirstName + " " + u.LastName + ":::" + date + ":::" + fp.Text + ":::" + Itoa(int(fp.Id))
		return commData
	}
}

// fonction Public
// permet de supprimer un commentaire sur une question
func ForumDelCommHandler(w http.ResponseWriter, r *http.Request) {
	var fp ForumPost
	fp.Id, _ = ParseInt(r.PostFormValue("id_commentaire"), 0, 64)
	fp.Delete()
	commData := "success"
	return commData
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("titre_post") == "" ||
		r.PostFormValue("categorie_post") == "" ||
		r.PostFormValue("contenu_post") == "" ||
		r.PostFormValue("user_id") == "" {
		return "error"
	} else {

		var f Forum

		isSolved, _ := ParseInt(r.PostFormValue("resolu_post"), 0, 64)
		idCat, _ := ParseInt(r.PostFormValue("categorie_post"), 0, 64)

		f.Title = r.PostFormValue("titre_post")
		f.ForumCategoryId = idCat
		f.IsSolved = isSolved
		f.Text = r.PostFormValue("contenu_post")
		f.IsOnline = 1
		f.UserId, _ = ParseInt(r.PostFormValue("user_id"), 0, 64)
		// si il y a pas encore un id met un 0
		f.Id, _ = ParseInt(r.PostFormValue("forum_id"), 0, 64)
		f.Id = f.Save()
		/*
			log.Println(f.Id)
			if f.Id == 0 {
				f.Id = f.Save()
			} else {
				f.Id = f.Update()
			}
		*/

		// String qui contient d'abord l'auteur du commentaire
		// puis son commentaire complet, séparés par ":::"
		commData := Itoa(int(f.Id)) + ":::" + "All good"
		return commData
	}
}
