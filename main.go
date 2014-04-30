package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/konginteractive/cme/app"
	"log"
	"net/http"

	// attention entre html/template et text/template le render est en autoescaping
	htmlTempl "html/template"
	textTempl "text/template"
)

// gestion des cessions
var store = sessions.NewCookieStore([]byte("something-very-secret"))

// gestion des templates
var templatesHtml *htmlTempl.Template
var templatesText *textTempl.Template

func init() {
	templatesHtml = htmlTempl.Must(htmlTempl.ParseGlob("./vues/*"))
	templatesText = textTempl.Must(textTempl.ParseGlob("./vues/*"))
}

func main() {

	r := mux.NewRouter()

	// listes des rootes
	r.HandleFunc("/", HomeHandler)

	// routages du forum
	r.HandleFunc("/forum", ForumHandler)

	r.HandleFunc("/forum/nouveau", ForumAddHandler)
	r.HandleFunc("/forum/search", ForumSearchHandler)
	r.HandleFunc("/forum/{category}", ForumCatHandler)
	r.HandleFunc("/forum/post/{id:[0-9]+}", ForumPostHandler)

	// routages des élèves
	r.HandleFunc("/eleves", StudentHandler)
	r.HandleFunc("/eleves/2014/henrilepic", StudentFicheHandler)

	// routage des tutoriels
	r.HandleFunc("/tutoriels", TutoHandler)
	r.HandleFunc("/tutoriel/nouveau", TutoAddHandler)
	r.HandleFunc("/tutoriel/search", TutoSearchHandler)
	r.HandleFunc("/tutoriel/{category}", TutoCatHandler)
	r.HandleFunc("/tutoriel/post/{id:[0-9]+}", TutoPostHandler)

	// routages des actualités
	r.HandleFunc("/actualites", NewsHandler)

	// rootage de la page connexion
	r.HandleFunc("/connexion", ConnexionHandler)

	r.HandleFunc("/valider-connexion", ConnexionPostHandler)

	r.HandleFunc("/get-connexion", ConnexionGet)

	r.HandleFunc("/inscription", InscriptionHandler)
	r.HandleFunc("/mon-compte", MonCompteHandler)

	// Rootage des pages d'informations
	r.HandleFunc("/qui-sommes-nous", QuiSommesNousHandler)
	r.HandleFunc("/pourquoi-une-association", PourquoiUneAssoHandler)

	//gestion des fichiers statiques
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}

// affichage de la home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	/*
		var ph PageHome
		Render(w, ph.View(), r)
	*/
}

// affichage du forum
func ForumHandler(w http.ResponseWriter, r *http.Request) {

	pf := new(PageForum)

	value := r.FormValue("p")
	if value == "" {
		pf.View()
	} else {
		pf.ViewPaged(value)
	}

	//insersion dans l'interface Page
	var p Page
	p = pf
	Render(w, p, r)

}

// affichage de la recherche dans le forum
func ForumSearchHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForum)
	q := r.FormValue("q")
	if q == "" {
		pf.View()
	} else {
		pf.ViewSearch(q)
	}

	//insersion dans l'interface Page
	var p Page
	p = pf
	Render(w, p, r)
}

// affichage du formulaire du forum
func ForumAddHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForum)
	f, v := pf.ValidateForm(r)
	if v {
		log.Print("VALIDE!!")
		f.Save()
	} else {
		log.Print("NON VALIDE!!")
	}
	//insersion dans l'interface Page
	var p Page
	p = pf
	Render(w, p, r)
}

// affichage d'une catégorie dans le forum
func ForumCatHandler(w http.ResponseWriter, r *http.Request) {
	pf := new(PageForum)

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	value := r.FormValue("p")

	if value == "" {
		pf.ViewCategory(category)
	} else {
		pf.ViewCategoryPaged(category, value)
	}
	//insersion dans l'interface Page
	var p Page
	p = pf
	Render(w, p, r)
}

// affichage d'une question du forum
func ForumPostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	pfp := new(PageForumPost)
	pfp.View(id)

	//insersion dans l'interface Page
	var p Page
	p = pfp
	Render(w, p, r)
}

// affichage de la liste des étudants
func StudentHandler(w http.ResponseWriter, r *http.Request) {

	pu := new(PageUser)
	pu.View()

	//insersion dans l'interface Page
	var p Page
	p = pu
	Render(w, p, r)
}

// affichage de la fiche étudiant
func StudentFicheHandler(w http.ResponseWriter, r *http.Request) {

	pu := new(PageUser)
	pu.ViewFiche()

	//insersion dans l'interface Page
	var p Page
	p = pu
	Render(w, p, r)
}

// affichage de la liste des tutos
func TutoHandler(w http.ResponseWriter, r *http.Request) {

	pt := new(PageTutorial)
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

// affichage du formulaire d'ajout d'un tuto
func TutoAddHandler(w http.ResponseWriter, r *http.Request) {

	pt := new(PageTutorial)

	t, v := pt.ValidateForm(r)

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
	Render(w, p, r)
}

// affichage d'une catégorie d'un tutoriel
func TutoCatHandler(w http.ResponseWriter, r *http.Request) {
	pt := new(PageTutorial)

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
	pt := new(PageTutorial)

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

// affichage d'un tuto
func TutoPostHandler(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		id := vars["id"]
	*/
	fp := new(PageForum)
	fp.View()

	//insersion dans l'interface Page
	var p Page
	p = fp
	Render(w, p, r)

}

// affichage de la liste des news
func NewsHandler(w http.ResponseWriter, r *http.Request) {

	pn := new(PageNews)
	pn.View()

	//insersion dans l'interface Page
	var p Page
	p = pn
	Render(w, p, r)
}

// affichage de la connexion
func ConnexionHandler(w http.ResponseWriter, r *http.Request) {

	pc := new(PageConnexion)
	pc.View()

	//insersion dans l'interface Page
	var p Page
	p = pc
	Render(w, p, r)

}

func ConnexionPostHandler(w http.ResponseWriter, r *http.Request) {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "cme_connecte")
	// Set some session values.
	session.Values["name"] = "Antoine"
	session.Values["id"] = 1
	session.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	// Save it.
	session.Save(r, w)
}

func ConnexionGet(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cme_connecte")

	fmt.Fprintf(w, "Hi there, I love %s!", session.Values["name"])

	// permet d'afficher quelque chose
	//var pc PageConnexion
	//Render(w, pc.View())
}

// affichage du formulaire d'inscription
func InscriptionHandler(w http.ResponseWriter, r *http.Request) {

	pi := new(PageInscription)
	pi.View()

	//insersion dans l'interface Page
	var p Page
	p = pi
	Render(w, p, r)
}

func MonCompteHandler(w http.ResponseWriter, r *http.Request) {
	var pc PageCompte
	Render(w, pc.View())
}

// affichage de la page qui sommes nous
func QuiSommesNousHandler(w http.ResponseWriter, r *http.Request) {

	pqsn := new(PageQuiSommesNous)
	pqsn.View()

	//insersion dans l'interface Page
	var p Page
	p = pqsn
	Render(w, p, r)
}

// affichage de la page pourquoi une association
func PourquoiUneAssoHandler(w http.ResponseWriter, r *http.Request) {

	ppua := new(PagePourquoiUneAsso)
	ppua.View()

	//insersion dans l'interface Page
	var p Page
	p = ppua
	Render(w, p, r)
}

// fonction permettant de rendu des pages en fonction du type HTML ou TXT
func Render(w http.ResponseWriter, p Page, r *http.Request) {

	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")

	session, _ := store.Get(r, "cme_connecte")
	log.Println(session.Values["id"])

	var u User
	//u.Id = int64(int(session.Values["id"]))
	//u.FirstName = string(session.Values["name"])
	u.Id = 2
	u.FirstName = "Antoine"
	p.SetSessionData(u)

	var err error
	// permet d'afficher avec le rendu html ou non
	if p.IsHtmlRender() == false {
		err = templatesHtml.ExecuteTemplate(w, Templ, p)
	} else {
		err = templatesText.ExecuteTemplate(w, Templ, p)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
