package app

import (
	"github.com/gorilla/sessions"
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

// fonction permettant de rendu des pages en fonction du type HTML ou TXT
func Render(w http.ResponseWriter, p Page, r *http.Request) {

	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "text/html")

	session, err := store.Get(r, "cme_connecte")

	if err == nil && session.Values["id"] != nil && session.Values["id"] != "0" {
		var u User
		u.Id = session.Values["id"].(int64)
		u.FirstName = session.Values["name"].(string)
		p.SetSessionData(u)
	} else {
		// permet de supprimer la session en cours
		session.Options.MaxAge = -1
		var u User
		p.SetSessionData(u)
	}

	//var err error
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
