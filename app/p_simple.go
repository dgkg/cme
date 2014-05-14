package app

import (
	"github.com/gorilla/mux"
	"net/http"
	. "strconv"
	"time"
)

// PageSimple est un type qui permet d'afficher des pages d'information
type PageSimple struct {
	Id        int64
	Url       string
	Title     string
	Text      string
	IsOnline  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	PageWeb   `sql:"-"` // Ignore this field
}

// affichage d'une question du forum
func PageSimpleHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	pageUrl := vars["pageUrl"]
	// surcharge de la variable d'affichage
	Templ = "simple"
	// initialisation de l'objet PageSimple
	var ps PageSimple
	ps.RenderHtml = true
	ps.Url = pageUrl
	ps.MainClass = "pageinfo"
	ps = ps.Get()
	//insersion dans l'interface Page
	var p Page
	p = ps
	Render(w, p, r)
}

// fonction public
// permet d'enregistrer les éléments du formulaire
func (ps PageSimple) Save() int64 {
	db.Save(&ps)
	if ps.Id != 0 {
		return ps.Id
	} else {
		db.Last(&ps)
		return ps.Id
	}
}

// fonction public
// permet de supprimer une page simple
func (ps PageSimple) Delete() {
	db.Delete(&ps)
}

// permet de récupérer une page simple
// en fonction de son id si il est setup ou de son url
func (ps PageSimple) Get() PageSimple {
	if ps.Id != 0 {
		db.Where("is_online = ? AND id = ?", "1", Itoa(int(ps.Id))).Find(&ps)
	} else {
		db.Where("is_online = ? AND url = ?", "1", ps.Url).Find(&ps)
	}
	return ps
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (ps PageSimple) IsHtmlRender() bool {
	return ps.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (ps PageSimple) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		ps.SessIdUser = u.Id
		ps.SessIsLogged = true
		ps.SessNameUser = u.FirstName
		v = true
	}
	return
}

/*







*/
