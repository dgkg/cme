package main

import (
	C "github.com/konginteractive/cme/controler"
	M "github.com/konginteractive/cme/model"
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

	// routages des élèves
	r.HandleFunc("/eleves", StudentHandler)

	// routage des tutoriels
	r.HandleFunc("/tutoriels", TutoHandler)
	r.HandleFunc("/tutoriels/nouveau", TutoAddHandler)
	r.HandleFunc("/actualites", NewsHandler)

	// routages des actualités
	r.HandleFunc("/actualites", NewsHandler)

	//gestion des fichiers statiques
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	Render(w, "index", C.HomeView())
}

func ForumHandler(w http.ResponseWriter, r *http.Request) {

	p := r.FormValue("p")
	if p == "" {
		Render(w, C.ForumTempl, C.ForumView())
	} else {
		Render(w, C.ForumTempl, C.ForumViewPaged(p))
	}

}

func ForumSearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		Render(w, C.ForumTempl, C.ForumView())
	} else {
		Render(w, C.ForumTempl, C.ForumViewSearch(q))
	}
}

func ForumAddHandler(w http.ResponseWriter, r *http.Request) {
	isValid := C.ValidateForum(r)

	if isValid {
		log.Print("VALIDE!!")
		C.SetPostForum(r)
	} else {
		log.Print("NON VALIDE!!")
	}

	Render(w, C.ForumTempl, C.ForumAddView())
}
func ForumCatHandler(w http.ResponseWriter, r *http.Request) {
	// récupère la catégorie sélectionnée
	vars := mux.Vars(r)
	category := vars["category"]
	// récupère la page en cours sélectionnée
	p := r.FormValue("p")

	if p == "" {
		Render(w, C.ForumTempl, C.FormViewCategory(category))
	} else {
		Render(w, C.ForumTempl, C.FormViewCategoryPaged(category, p))
	}
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, C.UserTempl, C.UserView())
}

func TutoHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, C.TutorialTempl, C.TutorialView())
}

func TutoAddHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, C.TutorialAddTempl, C.TutorialAddView())
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, C.NewsTempl, C.NewsView())
}

func Render(w http.ResponseWriter, tmpl string, p M.Page) {
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
