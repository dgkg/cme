package main

import (
	. "github.com/konginteractive/cme/app"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("./vues/*"))
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

	//gestion des fichiers statiques
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var ph PageHome
	Render(w, ph.View())
}

func ForumHandler(w http.ResponseWriter, r *http.Request) {
	var pf PageForum
	p := r.FormValue("p")
	if p == "" {
		Render(w, pf.View())
	} else {
		Render(w, pf.ViewPaged(p))
	}
}

func ForumSearchHandler(w http.ResponseWriter, r *http.Request) {
	var pf PageForum

	q := r.FormValue("q")
	if q == "" {
		Render(w, pf.View())
	} else {
		Render(w, pf.ViewSearch(q))
	}
}

func ForumAddHandler(w http.ResponseWriter, r *http.Request) {

	var pf PageForum

	f, v := pf.ValidateForm(r)

	log.Print("attentions : " + f.Title)

	if v {
		log.Print("VALIDE!!")
		f.Save()
	} else {
		log.Print("NON VALIDE!!")
	}

	Render(w, pf.ViewAdd())
}

func ForumCatHandler(w http.ResponseWriter, r *http.Request) {
	var pf PageForum

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	p := r.FormValue("p")

	if p == "" {
		Render(w, pf.ViewCategory(category))
	} else {
		Render(w, pf.ViewCategoryPaged(category, p))
	}
}

func ForumPostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var pfp PageForumPost
	Render(w, pfp.View(id))
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var pu PageUser
	Render(w, pu.View())
}

func StudentFicheHandler(w http.ResponseWriter, r *http.Request) {
	var pu PageUser
	Render(w, pu.ViewFiche())
}

func TutoHandler(w http.ResponseWriter, r *http.Request) {
	var pt PageTutorial
	p := r.FormValue("p")
	if p == "" {
		Render(w, pt.View())
	} else {
		Render(w, pt.ViewPaged(p))
	}
}

func TutoAddHandler(w http.ResponseWriter, r *http.Request) {

	var pt PageTutorial

	t, v := pt.ValidateForm(r)

	log.Print("attentions : " + t.Title)

	if v {
		log.Print("VALIDE!!")
		t.Save()
	} else {
		log.Print("NON VALIDE!!")
	}

	Render(w, pt.ViewAdd())

}

func TutoCatHandler(w http.ResponseWriter, r *http.Request) {

	var pt PageTutorial

	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	p := r.FormValue("p")

	if p == "" {
		Render(w, pt.ViewCategory(category))
	} else {
		Render(w, pt.ViewCategoryPaged(category, p))
	}
}

func TutoSearchHandler(w http.ResponseWriter, r *http.Request) {
	var pt PageTutorial

	q := r.FormValue("q")
	if q == "" {
		Render(w, pt.View())
	} else {
		Render(w, pt.ViewSearch(q))
	}
}

func TutoPostHandler(w http.ResponseWriter, r *http.Request) {
	//var fp PageForum
	//Render(w, C.ForumTempl, C.ForumViewPost())
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	var pn PageNews
	Render(w, pn.View())
}

func ConnexionHandler(w http.ResponseWriter, r *http.Request) {
	var pc PageConnexion
	Render(w, pc.View())
}

func Render(w http.ResponseWriter, p Page) {

	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, Templ, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
