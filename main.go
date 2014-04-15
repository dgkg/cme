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
	r.HandleFunc("/forum", ForumHandler)
	r.HandleFunc("/forum/add", ForumAddHandler)
	r.HandleFunc("/eleves", StudentHandler)
	r.HandleFunc("/tutoriels", TutoHandler)
	r.HandleFunc("/tutoriel/add", TutoAddHandler)
	r.HandleFunc("/actualites", NewsHandler)

	/*
		// Forum Handlers
		r.HandleFunc("/admin/forum", ForumHandler)
		r.HandleFunc("/admin/forum/add", ForumAddHandler)
		r.HandleFunc("/admin/forum/edit/{id:[0-9]+}", ForumEditHandler)
		r.HandleFunc("/admin/forum/edit/{id:[0-9]+}/post/", ForumPostHandler)
		r.HandleFunc("/admin/forum/edit/{id:[0-9]+}/post/add", ForumPostAddHandler)
		r.HandleFunc("/admin/forum/edit/{id:[0-9]+}/post/edit/{idpost:[0-9]+}", ForumPostEditHandler)
		r.HandleFunc("/admin/forum/category", ForumCategoryHandler)
		r.HandleFunc("/admin/forum/category/add", ForumCategoryAddHandler)
		r.HandleFunc("/admin/forum/category/edit/{id:[0-9]+}", ForumCategoryEditHandler)
		r.HandleFunc("/admin/forum/keyword", ForumKeywordHandler)
		r.HandleFunc("/admin/forum/keyword/add", ForumKeywordAddHandler)
		r.HandleFunc("/admin/forum/keyword/edit/{id:[0-9]+}", ForumKeywordEditHandler)

		// News Handlers
		r.HandleFunc("/admin/news", NewsHandler)
		r.HandleFunc("/admin/news/add", NewsAddHandler)
		r.HandleFunc("/admin/news/edit/{id:[0-9]+}", NewsEditHandler)
		r.HandleFunc("/admin/news/category", NewsCategoryHandler)
		r.HandleFunc("/admin/news/category/add", NewsCategoryAddHandler)
		r.HandleFunc("/admin/news/category/edit/{id:[0-9]+}", NewsCategoryEditHandler)
		r.HandleFunc("/admin/news/keyword", NewsKeywordHandler)
		r.HandleFunc("/admin/news/keyword/add", NewsKeywordAddHandler)
		r.HandleFunc("/admin/news/keyword/edit/{id:[0-9]+}", NewsKeywordEditHandler)

		// Tutorial Handlers
		r.HandleFunc("/admin/tutorial", TutorialHandler)
		r.HandleFunc("/admin/tutorial/add", TutorialAddHandler)
		r.HandleFunc("/admin/tutorial/edit/{id:[0-9]+}", TutorialEditHandler)
		r.HandleFunc("/admin/tutorial/category", TutorialCategoryHandler)
		r.HandleFunc("/admin/tutorial/category/add", TutorialCategoryAddHandler)
		r.HandleFunc("/admin/tutorial/category/edit/{id:[0-9]+}", TutorialCategoryEditHandler)
		r.HandleFunc("/admin/tutorial/keyword", TutorialKeywordHandler)
		r.HandleFunc("/admin/tutorial/keyword/add", TutorialKeywordAddHandler)
		r.HandleFunc("/admin/tutorial/keyword/edit/{id:[0-9]+}", TutorialKeywordEditHandler)

		// User Handlers
		r.HandleFunc("/admin/user", UserHandler)
		r.HandleFunc("/admin/user/add", UserAddHandler)
		r.HandleFunc("/admin/user/edit/{id:[0-9]+}", UserEditHandler)
		r.HandleFunc("/admin/user/edit/{id:[0-9]+}/media", UserMediaHandler)
		r.HandleFunc("/admin/user/edit/{id:[0-9]+}/media/add", UserMediaAddHandler)
		r.HandleFunc("/admin/user/edit/{id:[0-9]+}/media/edit/{idmedia:[0-9]+}", UserMediaEditHandler)
	*/
	//

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
	Render(w, C.ForumTempl, C.ForumView())
}

func ForumAddHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "forum_add", C.HomeView())
}
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, C.UserTempl, C.UserView())
}
func TutoHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, C.TutorialTempl, C.TutorialView())
}
func TutoAddHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "tutoriel_add", C.HomeView())
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
