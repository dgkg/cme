package main

import (
	"github.com/gorilla/mux"
	. "github.com/konginteractive/cme/app"
	"log"
	"net/http"
)

func main() {

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

	// routages des élèves
	r.HandleFunc("/eleves", StudentHandler)
	r.HandleFunc("/eleves/{year:[0-9]+}/{firstName:[a-z-]+}_{lastName:[a-z-]+}", StudentFicheHandler)
	//r.HandleFunc("/eleves/{year:[0-9]+}/{firstName:[a-z-]+}_{lastName:[a-z-]+}/{project:[a-z-]+}", ProjetHandler)

	// routage des tutoriels
	r.HandleFunc("/tutoriels", TutoHandler)
	r.HandleFunc("/tutoriel/nouveau", TutoAddHandler)
	r.HandleFunc("/tutoriel/search", TutoSearchHandler)
	r.HandleFunc("/tutoriel/{category}", TutoCatHandler)
	r.HandleFunc("/tutoriel/post/{id:[0-9]+}", TutoPostHandler)
	r.HandleFunc("/tutoriel/post/edit/{id:[0-9]+}", TutoEditHandler)

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
	r.HandleFunc("/forum/nouveau/submitform", SubmitFormHandler)
	r.HandleFunc("/forum/post/nouvcomm", ForumNouvCommHandler)
	r.HandleFunc("/forum/post/delcomm", ForumDelCommHandler)
	r.HandleFunc("/tutoriel/nouveau/submitform", SubmitTutorialHandler)
	r.HandleFunc("/tutoriel/post/nouvcomm", TutorialNouvCommHandler)
	r.HandleFunc("/tutoriel/post/delcomm", TutorialDelCommHandler)
	r.HandleFunc("/inscription/submitform", InscFormHandler)
	r.HandleFunc("/mon-compte/update", EditCompteHandler)
	r.HandleFunc("/mon-compte/upload", UploadHandler)
	r.HandleFunc("/ajax/userProject/getByIdCat/{id:[0-9]+}", UserProjectAjaxHandler)

	//r.HandleFunc("/ajax/connect", AjaxConnectHandler)
	r.HandleFunc("/get-projects", GetProjectsHandler)

	// Tests
	//r.HandleFunc("/test-post", TestPostHandler)

	// Page simple
	r.HandleFunc("/{pageUrl:[a-z-]+}", PageSimpleHandler)

	//gestion des fichiers statiques
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}
