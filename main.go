package main

import (
	"github.com/gorilla/mux"
	. "github.com/konginteractive/cme/app/handler"
	"log"
	"net/http"
)

func main() {

	// création du routeur
	r := mux.NewRouter()

	// listes des routes
	r.HandleFunc("/", HomeHandler)

	// routages du forum
	r.HandleFunc("/forum", ForumHandler)
	r.HandleFunc("/forum/nouveau", ForumAddHandler)
	r.HandleFunc("/forum/search", ForumSearchHandler)
	r.HandleFunc("/forum/{category}", ForumCatHandler)
	r.HandleFunc("/forum/post/{id:[0-9]+}", ForumPostHandler)
	r.HandleFunc("/forum/post/edit/{id:[0-9]+}", ForumEditHandler)
	// AJAX
	r.HandleFunc("/ajax/forum/add", SubmitFormHandler)
	r.HandleFunc("/ajax/forum-post/add", ForumNouvCommHandler)
	r.HandleFunc("/ajax/forum-post/del", ForumDelCommHandler)

	// routages des élèves
	r.HandleFunc("/eleves", StudentHandler)
	r.HandleFunc("/eleves/search", StudentSearchHandler)
	r.HandleFunc("/eleves/{year:[0-9]+}/{firstName:[a-z-]+}_{lastName:[a-z-]+}", StudentFicheHandler)
	//r.HandleFunc("/eleves/{year:[0-9]+}/{firstName:[a-z-]+}_{lastName:[a-z-]+}/{project:[a-z-]+}", ProjetHandler)

	// routage des tutoriels
	r.HandleFunc("/tutoriels", TutoHandler)
	r.HandleFunc("/tutoriel/nouveau", TutoAddHandler)
	r.HandleFunc("/tutoriel/search", TutoSearchHandler)
	r.HandleFunc("/tutoriel/{category}", TutoCatHandler)
	r.HandleFunc("/tutoriel/post/{id:[0-9]+}", TutoPostHandler)
	r.HandleFunc("/tutoriel/post/edit/{id:[0-9]+}", TutoEditHandler)
	// AJAX
	r.HandleFunc("/ajax/tutoriel/add", SubmitTutorialHandler)
	r.HandleFunc("/ajax/tutoriel-post/add", TutorialNouvCommHandler)
	r.HandleFunc("/ajax/tutoriel-post/del", TutorialDelCommHandler)

	// routages des actualités
	r.HandleFunc("/actualites", NewsHandler)
	r.HandleFunc("/actualite/{name:[a-z-]+}", NewsUniqueHandler)

	// rootage de la page connexion
	r.HandleFunc("/connexion", ConnexionHandler)
	r.HandleFunc("/deconnexion", DeconnexionHandler)
	r.HandleFunc("/valider-connexion", ConnexionPostHandler)
	r.HandleFunc("/get-connexion", ConnexionGet)
	r.HandleFunc("/inscription", InscriptionHandler)
	r.HandleFunc("/mon-compte", MonCompteHandler)
	// AJAX
	r.HandleFunc("/ajax/inscription/submitform", InscFormHandler)
	r.HandleFunc("/ajax/mon-compte/update", EditCompteHandler)
	r.HandleFunc("/ajax/mon-compte/avatar", AvatarHandler)
	r.HandleFunc("/ajax/mon-compte/cover", CoverHandler)
	r.HandleFunc("/ajax/userProject/getByIdCat/{id:[0-9]+}", UserProjectAjaxHandler)
	r.HandleFunc("/ajax/userProject/getByIdCat/{id:[0-9]+}/{page:[0-9]+}", UserProjectMoreInCatAjaxHandler)
	r.HandleFunc("/ajax/userProject/getMore/{page:[0-9]+}", UserProjectMoreAjaxHandler)
	r.HandleFunc("/ajax/up/create", UPCreateAjaxHandler)
	r.HandleFunc("/ajax/upi/create", UPICreateAjaxHandler)
	r.HandleFunc("/ajax/get-projects", GetProjectsHandler)

	// Page simple
	r.HandleFunc("/page/{pageUrl:[a-z-]+}", PageSimpleHandler)

	//gestion des fichiers statiques
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// insersion des routes dans l'objet http
	// et lancement du serveur
	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}
