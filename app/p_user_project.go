package app

import (
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/kennygrant/sanitize"
	//"log"
	"net/http"
	. "strconv"
	//"time"
	"encoding/json"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
)

type PageUserProject struct {
	User
	PagesList []Paginate
	PageWeb
}

// Public function
// permet de récupérer une liste de projets
func UserProjectAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	//page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCat(3)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}
		fmt.Fprint(w, string(b))
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
func UserProjectMoreAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if page == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		ups := up.getListPaged(page, 3)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}
		fmt.Fprint(w, string(b))
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
// et avec une catégorie sélectionnée
func UserProjectMoreInCatAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 || page == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCatPaged(page, 3)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}
		fmt.Fprint(w, string(b))
	}
}

// permet de créer un userProject
func UserProjectCreateAjaxHandler(w http.ResponseWriter, r *http.Request) {

	// création d'un user project
	var up UserProject
	up.UserId, _ = ParseInt(r.PostFormValue("id_user"), 0, 64)
	up.Title = r.PostFormValue("title")
	up.Description = r.PostFormValue("Description")
	up.IsOnline = 1 // permet qu'il soit en ligne
	up.Id = up.Save()

	b, err := json.Marshal(up)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	}
	fmt.Fprint(w, string(b))
}

// Upload de fichiers avec Go
// Code original : https://gist.github.com/sanatgersappa/5127317#file-app-go
func UserProjectUploadAjaxHandler(w http.ResponseWriter, r *http.Request) {

	// création d'une image dans user project
	var upi UserProjectImage
	// récupération de la variable id
	upi.ProjectId, _ = ParseInt(r.PostFormValue("id_user_project"), 0, 64)
	upi.Title = r.PostFormValue("title")
	upi.Description = r.PostFormValue("Description")

	var strNomFichier string

	//parse the multipart form in the request
	err := r.ParseMultipartForm(100000)
	if err != nil {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
		/*
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		*/
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["myfiles"]
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Création d'un nom sanitizé pour
		strNomFichier = sanitize.Name(files[i].Filename)

		// Validation des extensions de fichiers
		if filepath.Ext(strNomFichier) != ".jpg" &&
			filepath.Ext(strNomFichier) != ".jpeg" &&
			filepath.Ext(strNomFichier) != ".png" &&
			filepath.Ext(strNomFichier) != ".gif" {
			return
		}

		//create destination file making sure the path is writeable.
		urlImage := "./public/img/uploads/users/" + userId + "/" + files[i].Filename
		dst, err := os.Create(urlImage)
		//defer dst.Close()
		if err != nil {
			log.Println(err)
		}

		// Envoi du nom du fichier pour l'ajout à la BD
		upi.Url = urlImage
		upi.Id = upi.Save()
		up.Images = upi

		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	/*
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
	*/
}
