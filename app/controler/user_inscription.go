package app

import (
	"fmt"
	"github.com/gorilla/sessions"
	. "github.com/konginteractive/cme/app/model"
	"io"
	"log"
	"net/http"
	"os"
	. "strconv"
)

type PageInscription struct {
	PageWeb
}

// affichage du formulaire d'inscription
func InscriptionHandler(w http.ResponseWriter, r *http.Request) {

	pi := new(PageInscription)
	idUser, isValid := pi.ValidateDataInscrption(r)

	var u User
	u.Id = idUser
	u = u.getById()

	if isValid {
		session, _ := store.Get(r, "cme_connecte")
		// Set some session values.
		session.Values["id"] = u.Id
		session.Values["name"] = u.FirstName
		session.Options = &sessions.Options{
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		// Save it.
		session.Save(r, w)
		// Set data into current page
		pi.SessIdUser = u.Id
		pi.SessNameUser = u.FirstName
		pi.SessIsLogged = true
	}

	pi.View()

	//insersion dans l'interface Page
	var p Page
	p = pi
	Render(w, p, r)
}

func (pi *PageInscription) ValidateDataInscrption(r *http.Request) (idUser int64, isValid bool) {

	//log.Println("Validation Data Inscriptin applé")
	//log.Println(r.PostFormValue("prenom"))
	//log.Println(r.PostFormValue("nom"))
	//log.Println(r.PostFormValue("email"))
	//log.Println(r.PostFormValue("pass"))
	//log.Println("Fin Validation Data Inscriptin applé")

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("prenom") == "" ||
		r.PostFormValue("nom") == "" ||
		r.PostFormValue("email") == "" ||
		r.PostFormValue("pass") == "" {

		idUser = 0
		isValid = false
		return
	} else {

		var u User
		u.FirstName = r.PostFormValue("prenom")
		u.LastName = r.PostFormValue("nom")
		u.Email = r.PostFormValue("email")
		u.Pass = r.PostFormValue("pass")
		u.IsOnline = 1
		u.PhotoProfil = "user_inconnu.jpg"
		u.PhotoCover = "projet.jpg"
		idUser = u.Save()

		// Création d'un répertoire utilisateur
		urlPath := "./public/img/uploads/users/" + Itoa(int(idUser))
		isExist, _ := exists(urlPath)
		if !isExist {
			log.Println(urlPath)
			os.MkdirAll(urlPath, 0700)
		}

		// Copie de l'image user par défault
		userPath := urlPath + "/user_inconnu.jpg"
		cp(userPath, "./public/img/ui/user_inconnu.jpg")

		projetPath := urlPath + "/projet.jpg"
		cp(projetPath, "./public/img/ui/projet.jpg")

		isValid = true
		return
	}

}

// Vérifie sur le path fourni existe ou non
// retourne true sur le path existe
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Copie un fichier
// Source : https://gist.github.com/elazarl/5507969
func cp(dst, src string) error {
	s, err := os.Open(src)
	log.Println("s = ")
	log.Println(s)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	log.Println("d = ")
	log.Println(d)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func InscFormHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if r.PostFormValue("prenom") == "" ||
		r.PostFormValue("nom") == "" ||
		r.PostFormValue("email") == "" ||
		r.PostFormValue("pass") == "" {

		fmt.Fprint(w, "error")
	} else {

		var u User

		u.FirstName = r.PostFormValue("prenom")
		u.LastName = r.PostFormValue("nom")
		u.Email = r.PostFormValue("email")
		u.Pass = r.PostFormValue("pass")
		u.IsOnline = 1
		u.Save()
	}
}

// fonction pour permettre de créer une page
func CreatePageInscription() *PageInscription {
	return new(PageInscription)
}

func (pi *PageInscription) View() {

	log.Println("Inscription appelé")

	Templ = "inscription"

	pi.Title = "Inscription"
	pi.MainClass = "inscription"
	pi.RenderHtml = false
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageInscription) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageInscription) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
	return
}
