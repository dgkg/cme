package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/model"
	htmlTempl "html/template"
	"net/http"
	textTempl "text/template"
)

// gestion des cessions
var store = sessions.NewCookieStore([]byte("something-very-secret"))

// gestion des templates
var templatesHtml *htmlTempl.Template
var templatesText *textTempl.Template

func init() {
	// /Users/henrilepic/gocode/src/github.com/konginteractive/cme/
	templatesHtml = htmlTempl.Must(htmlTempl.ParseGlob("./app/vues/*"))
	templatesText = textTempl.Must(textTempl.ParseGlob("./app/vues/*"))
	// permet d'avoir quelques variables
	fmt.Println("YS : " + YS + " / MS : " + MS)
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

// permet de renvoyer un objet générique en json
func RenderJson(w http.ResponseWriter, msg interface{}) {

	b, err := json.Marshal(msg)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, string(b))
}
