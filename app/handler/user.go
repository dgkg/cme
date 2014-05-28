package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/controler"
	. "github.com/konginteractive/cme/app/model"
	"log"
	"net/http"
	. "strconv"
)

// affichage de la fiche étudiant
func (h *Handlers) StudentFicheHandler() (p Page) {

	pu := new(PageUser)
	//var u User
	// récupération de la variable de l'utilisateur
	vars := mux.Vars(h.R)
	pu.User.Graduation = vars["year"]
	pu.User.FirstName = vars["firstName"]
	pu.User.LastName = vars["lastName"]

	isFound := pu.FindUser()
	if isFound {
		pu.View()
	} else {
		// todo réaliser une redirection vers page non trouvée
	}

	//insersion dans l'interface Page
	p = pu
	return p
}

// affichage de la connexion
func (h *Handlers) ConnexionHandler() (p Page) {

	var u User
	u.Email = h.R.PostFormValue("login")
	u.Pass = h.R.PostFormValue("pass")

	var connected bool
	if u.Email != "" && u.Pass != "" {
		connected, _ = u.LoginPassExist()
	}

	pc := new(PageConnexion)
	pc.View()

	if connected {
		session, _ := store.Get(h.R, "cme_connecte")
		// Set some session values.
		session.Values["id"] = u.Id
		session.Values["name"] = u.FirstName
		session.Options = &sessions.Options{
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		// Save it.
		session.Save(h.R, h.W)
		// Set data into current page
		pc.SessIdUser = u.Id
		pc.SessNameUser = u.FirstName
		pc.SessIsLogged = true
		//http.Redirect(w, r, "http://www.google.com", 301)
	}

	//insersion dans l'interface Page
	p = pc
	return p
}

// DeconnexionHandler permet de supprimer la connexion "cme_connecte"
// en changant la date d'expiration à -1
// puis affiche la home
func (h *Handlers) DeconnexionHandler() (p Page) {
	// récupère la cession en cours
	session, _ := store.Get(h.R, "cme_connecte")
	// modifie sa date d'expiration
	session.Options.MaxAge = -1
	// sauvegarde
	session.Save(h.R, h.W)
	// creation de l'objet PageSimple
	var ps PageSimple
	ps.SessIdUser = 0
	ps.SessIsLogged = false
	ps.SessNameUser = ""
	// surcharge du template à utiliser
	Templ = "deconnexion"
	//insersion dans l'interface Page
	p = ps
	return p
}

func (h *Handlers) ConnexionPostHandler() {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(h.R, "cme_connecte")
	// Set some session values.
	session.Values["name"] = "Antoine"
	session.Values["id"] = 1
	session.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	// Save it.
	session.Save(h.R, h.W)
}

// affichage du formulaire d'inscription
func (h *Handlers) InscriptionHandler() (p Page) {

	pi := new(PageInscription)
	idUser, isValid := pi.ValidateDataInscrption(h.R)

	var u User
	u.Id = idUser
	u = u.getById()

	if isValid {
		session, _ := store.Get(h.R, "cme_connecte")
		// Set some session values.
		session.Values["id"] = u.Id
		session.Values["name"] = u.FirstName
		session.Options = &sessions.Options{
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		// Save it.
		session.Save(h.R, h.W)
		// Set data into current page
		pi.SessIdUser = u.Id
		pi.SessNameUser = u.FirstName
		pi.SessIsLogged = true
	}

	pi.View()

	//insersion dans l'interface Page
	p = pi
	return p
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func (h *Handlers) InscFormHandler() (m string) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	if h.R.PostFormValue("prenom") == "" ||
		h.R.PostFormValue("nom") == "" ||
		h.R.PostFormValue("email") == "" ||
		h.R.PostFormValue("pass") == "" {

		return "error"
	} else {
		var u User
		u.FirstName = h.R.PostFormValue("prenom")
		u.LastName = h.R.PostFormValue("nom")
		u.Email = h.R.PostFormValue("email")
		u.Pass = h.R.PostFormValue("pass")
		u.IsOnline = 1
		u.Save()
	}
}

// affichage de la liste des étudants
func (h *Handlers) StudentHandler() (p Page) {
	pul := new(PageUserList)
	pul.View()
	p = pul
	return p
}

// affichage de la recherche dans la liste des étudants
func (h *Handlers) StudentSearchHandler() (p Page) {
	pul := new(PageUserList)
	q := h.R.FormValue("q")
	if q == "" {
		pul.View()
	} else {
		pul.ViewSearch(q)
	}
	p = pul
	return p
}

func (h *Handlers) MonCompteHandler() (p Page) {

	pc := new(PageCompte)

	session, err := store.Get(h.R, "cme_connecte")

	var u User

	if err == nil && session.Values["id"] != nil {
		u.Id = session.Values["id"].(int64)
	}

	pc.View(u.Id)

	//insersion dans l'interface Page
	p = pc
	return p
}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func (h *Handlers) EditCompteHandler() (m string) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	sectionRecue := h.R.PostFormValue("section")

	if sectionRecue == "" {

		fmt.Fprint(h.W, "error")

	} else {

		var u User

		u.Id, _ = ParseInt(h.R.PostFormValue("idUser"), 0, 64)
		//log.Println(u.Id)

		switch sectionRecue {

		case "savePhoto":
			//	@todo : Recevoir la photo +
			//	l'héberger dans public/img/ +
			//	retourner une confirmation

		case "saveBio":
			u.Text = h.R.PostFormValue("biographie")
			u.SaveBio()

		case "saveSocial":
			u.Facebook = h.R.PostFormValue("facebook")
			u.Twitter = h.R.PostFormValue("twitter")
			u.LinkedIn = h.R.PostFormValue("linkedin")
			u.SaveSocial()

		case "saveGraduation":
			u.Graduation = h.R.PostFormValue("graduation")
			u.SaveGrad()

		case "saveNomUtilisateur":
			u.FirstName = h.R.PostFormValue("prenom")
			u.LastName = h.R.PostFormValue("nom")
			u.SaveUserName()

		case "supprimerCompte":
			u.DeleteAccount()
		}
	}
}

// Upload de fichiers avec Go
// Code original : https://gist.github.com/sanatgersappa/5127317#file-app-go
func (h *Handlers) AvatarHandler() {

	// Initialisation des variables utiles
	var u User
	session, _ := store.Get(h.R, "cme_connecte")
	u.Id = session.Values["id"].(int64)
	userId := Itoa(int(u.Id))
	var strNomFichier string

	//parse the multipart form in the request
	err := h.R.ParseMultipartForm(100000)
	if err != nil {
		http.Error(h.W, err.Error(), http.StatusInternalServerError)
		return
	}

	//get a ref to the parsed multipart form
	m := h.R.MultipartForm

	//get the *fileheaders
	files := m.File["photo-upload"]

	file, err := files[0].Open()

	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Création d'un nom sanitizé pour
	strNomFichier = sanitize.Name(files[0].Filename)

	// Validation des extensions de fichiers
	if filepath.Ext(strNomFichier) != ".jpg" &&
		filepath.Ext(strNomFichier) != ".jpeg" &&
		filepath.Ext(strNomFichier) != ".png" &&
		filepath.Ext(strNomFichier) != ".gif" {
		return
	}

	//create destination file making sure the path is writeable.
	dst, err := os.Create("./public/img/uploads/users/" + userId + "/" + strNomFichier)

	// Vérification de l'existence du répertoire cible
	if err != nil {
		log.Println(err)
	}

	// Envoi du nom du fichier pour l'ajout à la BD
	u.SavePhoto(strNomFichier)

	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redimensionnement de l'image

	// open "test.jpg"
	rfile, err := os.Open("./public/img/uploads/users/" + userId + "/" + strNomFichier)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	rimg, err := jpeg.Decode(rfile)
	if err != nil {
		log.Fatal(err)
	}
	rfile.Close()

	image := resize.Thumbnail(300, 300, rimg, resize.Bilinear)

	out, err := os.Create("./public/img/uploads/users/" + userId + "/" + strNomFichier)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, image, nil)

}

// Upload de fichiers avec Go
// Code original : https://gist.github.com/sanatgersappa/5127317#file-app-go
func (h *Handlers) CoverHandler() {
	log.Println("CoverHandler appelé!")

	// Initialisation des variables utiles
	var u User
	session, _ := store.Get(r, "cme_connecte")
	u.Id = session.Values["id"].(int64)
	userId := Itoa(int(u.Id))
	var strNomFichier string

	//parse the multipart form in the request
	err := r.ParseMultipartForm(100000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["cover-upload"]

	file, err := files[0].Open()

	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Création d'un nom sanitizé pour
	strNomFichier = sanitize.Name(files[0].Filename)

	// Validation des extensions de fichiers
	if filepath.Ext(strNomFichier) != ".jpg" &&
		filepath.Ext(strNomFichier) != ".jpeg" &&
		filepath.Ext(strNomFichier) != ".png" &&
		filepath.Ext(strNomFichier) != ".gif" {
		return
	}

	//create destination file making sure the path is writeable.
	dst, err := os.Create("./public/img/uploads/users/" + userId + "/" + strNomFichier)

	// Vérification de l'existence du répertoire cible
	if err != nil {
		log.Println(err)
	}

	// Envoi du nom du fichier pour l'ajout à la BD
	u.SavePhoto(strNomFichier)

	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redimensionnement de l'image

	// open "test.jpg"
	rfile, err := os.Open("./public/img/uploads/users/" + userId + "/" + strNomFichier)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	rimg, err := jpeg.Decode(rfile)
	if err != nil {
		log.Fatal(err)
	}
	rfile.Close()

	image := resize.Thumbnail(1280, 200, rimg, resize.Bilinear)

	out, err := os.Create("./public/img/uploads/users/" + userId + "/" + strNomFichier)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, image, nil)

}
