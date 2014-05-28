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

type Handlers struct {
	W http.ResponseWriter
	R *http.Request
}

// affichage d'une question du forum
func (h *Handlers) ForumPostHandler() (p Page) {
	// récupération de la variable id
	vars := mux.Vars(h.R)
	id := vars["id"]
	// initialisation de l'objet PageForum
	pf := new(PageForum)
	pf.Forum.Id, _ = ParseInt(id, 0, 64)
	pf.View()
	//insersion dans l'interface Page
	p = pf
	return p
}

// affichage du formulaire du forum
func (h *Handlers) ForumAddHandler() (p Page) {
	pf := new(PageForum)
	pf.ViewAdd()
	//insersion dans l'interface Page
	p = pf
	return p
}

func (h *Handlers) ForumEditHandler() (p Page) {

	// récupère la catégorie sélectionnée
	vars := mux.Vars(h.R)
	id := vars["id"]

	pf := new(PageForum)
	pf.ViewEdit(id)

	//insersion dans l'interface Page
	p = pf
	return p
}

// affichage du forum
func (h *Handlers) ForumHandler() (p Page) {

	pf := new(PageForumList)

	value := h.R.FormValue("p")
	if value == "" {
		pf.View()
	} else {
		pf.ViewPaged(value)
	}

	//insersion dans l'interface Page
	p = pf
	return p
}

// affichage de la recherche dans le forum
func (h *Handlers) ForumSearchHandler() (p Page) {
	pf := new(PageForumList)
	q := h.R.FormValue("q")
	if q == "" {
		pf.View()
	} else {
		pf.ViewSearch(q)
	}

	//insersion dans l'interface Page
	p = pf
	return p
}

// affichage d'une catégorie dans le forum
func (h *Handlers) ForumCatHandler() (p Page) {
	pf := new(PageForumList)

	// récupère la catégorie sélectionnée
	vars := mux.Vars(h.R)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	value := h.R.FormValue("p")

	if value == "" {
		pf.ViewCategory(category)
	} else {
		pf.ViewCategoryPaged(category, value)
	}
	//insersion dans l'interface Page
	p = pf
	return p
}

// Public function
// permet d'ajouter un commenaire sur une fonction
func (h *Handlers) ForumNouvCommHandler() (m string) {
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if h.R.PostFormValue("val_commentaire") == "" ||
		h.R.PostFormValue("val_post_id") == "" ||
		h.R.PostFormValue("val_auteur_id") == "" ||
		h.R.PostFormValue("val_auteur_id") == "0" {
		// envoie un message d'erreur
		return "error"
	} else {
		// initialise l'objet ForumPost et récupère les données du formulaire
		var fp ForumPost
		fp.ForumId, _ = ParseInt(h.R.PostFormValue("val_post_id"), 0, 64)
		fp.UserId, _ = ParseInt(h.R.PostFormValue("val_auteur_id"), 0, 64)
		fp.Text = sanitize.HTML(h.R.PostFormValue("val_commentaire"))
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
func (h *Handlers) ForumDelCommHandler() (m string) {
	var fp ForumPost
	fp.Id, _ = ParseInt(h.R.PostFormValue("id_commentaire"), 0, 64)
	fp.Delete()
	commData := "success"
	return commData
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func (h *Handlers) SubmitFormHandler() (m string) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if h.R.PostFormValue("titre_post") == "" ||
		h.R.PostFormValue("categorie_post") == "" ||
		h.R.PostFormValue("contenu_post") == "" ||
		h.R.PostFormValue("user_id") == "" {
		return "error"
	} else {

		var f Forum

		isSolved, _ := ParseInt(h.R.PostFormValue("resolu_post"), 0, 64)
		idCat, _ := ParseInt(h.R.PostFormValue("categorie_post"), 0, 64)

		f.Title = h.R.PostFormValue("titre_post")
		f.ForumCategoryId = idCat
		f.IsSolved = isSolved
		f.Text = h.R.PostFormValue("contenu_post")
		f.IsOnline = 1
		f.UserId, _ = ParseInt(h.R.PostFormValue("user_id"), 0, 64)
		// si il y a pas encore un id met un 0
		f.Id, _ = ParseInt(h.R.PostFormValue("forum_id"), 0, 64)
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
