package app

import (
	"fmt"
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/model"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	. "strconv"
)

type PageCompte struct {
	PageWeb
	User
	Graduations []string
}

func MonCompteHandler(w http.ResponseWriter, r *http.Request) {

	pc := new(PageCompte)

	session, err := store.Get(r, "cme_connecte")

	var u User

	if err == nil && session.Values["id"] != nil {
		u.Id = session.Values["id"].(int64)
	}

	pc.View(u.Id)

	//insersion dans l'interface Page
	var p Page
	p = pc
	Render(w, p, r)

}

// Réception du POST envoyé en AJAX et ajout des
// données dans la BD
func EditCompteHandler(w http.ResponseWriter, r *http.Request) {

	// Validation des données
	// Si un des variables est vide, la func retourne un "error"
	// ce qui fait afficher une message d'erreur
	sectionRecue := r.PostFormValue("section")

	if sectionRecue == "" {

		fmt.Fprint(w, "error")

	} else {

		var u User

		u.Id, _ = ParseInt(r.PostFormValue("idUser"), 0, 64)
		//log.Println(u.Id)

		switch sectionRecue {

		case "savePhoto":
			//	@todo : Recevoir la photo +
			//	l'héberger dans public/img/ +
			//	retourner une confirmation

		case "saveBio":
			u.Text = r.PostFormValue("biographie")
			u.SaveBio()

		case "saveSocial":
			u.Facebook = r.PostFormValue("facebook")
			u.Twitter = r.PostFormValue("twitter")
			u.LinkedIn = r.PostFormValue("linkedin")
			u.SaveSocial()

		case "saveGraduation":
			u.Graduation = r.PostFormValue("graduation")
			u.SaveGrad()

		case "saveNomUtilisateur":
			u.FirstName = r.PostFormValue("prenom")
			u.LastName = r.PostFormValue("nom")
			u.SaveUserName()

		case "supprimerCompte":
			u.DeleteAccount()
		}
	}
}

// Upload de fichiers avec Go
// Code original : https://gist.github.com/sanatgersappa/5127317#file-app-go
func AvatarHandler(w http.ResponseWriter, r *http.Request) {

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
func CoverHandler(w http.ResponseWriter, r *http.Request) {
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

func (pc *PageCompte) View(id int64) {

	log.Println("Mon Compte appelé")

	Templ = "user_compte"

	pc.Title = "Votre compte"
	pc.MainClass = "user_compte"

	var u User
	u.Id = id
	u = u.getById()

	pc.Text = u.Text
	pc.Facebook = u.Facebook
	pc.Twitter = u.Twitter
	pc.LinkedIn = u.LinkedIn
	pc.Graduation = u.Graduation
	pc.FirstName = u.FirstName
	pc.LastName = u.LastName
	pc.Id = u.Id
	pc.PhotoCover = u.PhotoCover

	pc.Graduations = make([]string, 20)

	var anneeGrad int = 2014

	for key, _ := range pc.Graduations {
		pc.Graduations[key] = Itoa(anneeGrad)
		anneeGrad--
	}

	pc.RenderHtml = true

}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageCompte) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageCompte) SetSessionData(u User) (v bool) {
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
