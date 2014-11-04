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
	s := r.Host("la-communaute-de-maryse-eloy.com").Subrouter()
        //r.Host("{subdomain:[a-z]+}.la-communaute-de-maryse-eloy.com")
	// listes des routes
	s.HandleFunc("/", HomeHandler)

	// routages du forum
	s.HandleFunc("/forum", ForumHandler)
	s.HandleFunc("/forum/nouveau", ForumAddHandler)
	s.HandleFunc("/forum/search", ForumSearchHandler)
	s.HandleFunc("/forum/{category}", ForumCatHandler)
	s.HandleFunc("/forum/post/{id:[0-9]+}", ForumPostHandler)
	s.HandleFunc("/forum/post/edit/{id:[0-9]+}", ForumEditHandler)
	// AJAX
	s.HandleFunc("/ajax/forum/add", SubmitFormHandler)
	s.HandleFunc("/ajax/forum-post/add", ForumNouvCommHandler)
	s.HandleFunc("/ajax/forum-post/del", ForumDelCommHandler)

	// routages des élèves
	s.HandleFunc("/eleves", StudentHandler)
	s.HandleFunc("/eleves/search", StudentSearchHandler)
	s.HandleFunc("/eleves/{year:[0-9]+}/{firstName:[a-z-]+}_{lastName:[a-z-]+}", StudentFicheHandler)
	//r.HandleFunc("/eleves/{year:[0-9]+}/{firstName:[a-z-]+}_{lastName:[a-z-]+}/{project:[a-z-]+}", ProjetHandler)

	// routage des tutoriels
	s.HandleFunc("/tutoriels", TutoHandler)
	s.HandleFunc("/tutoriel/nouveau", TutoAddHandler)
	s.HandleFunc("/tutoriel/search", TutoSearchHandler)
	s.HandleFunc("/tutoriel/{category}", TutoCatHandler)
	s.HandleFunc("/tutoriel/post/{id:[0-9]+}", TutoPostHandler)
	s.HandleFunc("/tutoriel/post/edit/{id:[0-9]+}", TutoEditHandler)
	// AJAX
	s.HandleFunc("/ajax/tutoriel/add", SubmitTutorialHandler)
	s.HandleFunc("/ajax/tutoriel-post/add", TutorialNouvCommHandler)
	s.HandleFunc("/ajax/tutoriel-post/del", TutorialDelCommHandler)

	// routages des actualités
	s.HandleFunc("/actualites", NewsHandler)
	s.HandleFunc("/actualite/{name:[a-z-]+}", NewsUniqueHandler)

	// rootage de la page connexion
	s.HandleFunc("/connexion", ConnexionHandler)
	s.HandleFunc("/deconnexion", DeconnexionHandler)
	s.HandleFunc("/valider-connexion", ConnexionPostHandler)
	//r.HandleFunc("/get-connexion", ConnexionGet)
	s.HandleFunc("/inscription", InscriptionHandler)
	s.HandleFunc("/mon-compte", MonCompteHandler)
	// AJAX
	s.HandleFunc("/ajax/inscription/submitform", InscFormHandler)
	s.HandleFunc("/ajax/mon-compte/update", EditCompteHandler)
	s.HandleFunc("/ajax/mon-compte/avatar", AvatarHandler)
	s.HandleFunc("/ajax/mon-compte/cover", CoverHandler)
	s.HandleFunc("/ajax/userProject/getByIdCat/{id:[0-9]+}", UserProjectAjaxHandler)
	s.HandleFunc("/ajax/userProject/getByIdCat/{id:[0-9]+}/{page:[0-9]+}", UserProjectMoreInCatAjaxHandler)
	s.HandleFunc("/ajax/userProject/getMore/{page:[0-9]+}", UserProjectMoreAjaxHandler)
	s.HandleFunc("/ajax/up/create", UPCreateAjaxHandler)
	s.HandleFunc("/ajax/upi/create", UPICreateAjaxHandler)
	s.HandleFunc("/ajax/get-projects", GetProjectsHandler)

	// Page simple
	s.HandleFunc("/page/{pageUrl:[a-z-]+}", PageSimpleHandler)

	//gestion des fichiers statiques
	s.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// insersion des routes dans l'objet http
	// et lancement du serveur
	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)

}
