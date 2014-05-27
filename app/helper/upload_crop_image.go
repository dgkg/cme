//
// Karine Duroux : LCL à contacter pour l'ordre de virement
//
package helper

import (
	//"fmt"
	"github.com/kennygrant/sanitize"
	//"log"
	"net/http"
	//. "strconv"
	"errors"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
)

// uploadImage permet de charger une image sur le serveur depuis un formulaire HTML
// le formulaire doit avoir une balise avec :
// <form enctype="multipart/form-data" action="http://localhost:8080/upload" method="post">
//     <input type="file" name="image" />
//     <input type="submit" value="upload" />
// </form>
func uploadImage(destination string, r *http.Request) (imgPath string, err error) {
	var strNomFichier string

	//parse the multipart form in the request
	err = r.ParseMultipartForm(MAX_SIZE_FILE_UPLOAD)
	if err != nil {
		err = errors.New("le fichier n' a pas les droits pour être créé")
		return "", err
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["image"]
	file, err := files[0].Open()
	defer file.Close()
	if err != nil {
		err = errors.New("le fichier n' a pas les droits pour être ouvert")
		return "", err
	}
	// Création d'un nom sanitizé pour
	strNomFichier = sanitize.Name(files[0].Filename)
	// Validation des extensions de fichiers
	if filepath.Ext(strNomFichier) != ".jpg" &&
		filepath.Ext(strNomFichier) != ".jpeg" &&
		filepath.Ext(strNomFichier) != ".png" &&
		filepath.Ext(strNomFichier) != ".gif" {
		err = errors.New("le fichier a une extention qui n'est pas autorisé")
		return "", err
	}

	// permet de voir si le dossier année existe
	// si il n'existe pas le créer
	// si il n'y arrive pas envoyer une erreur
	destination = destination + YS + "/"
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err := os.Mkdir(destination, 0777)
		if err != nil {
			err = errors.New("le dossier n'arrive pas à se créer")
			return "", err
		}
	}

	// permet de voir si le dossier mois existe
	// si il n'existe pas le créer
	// si il n'y arrive pas envoyer une erreur
	destination = destination + MS + "/"
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err := os.Mkdir(destination, 0777)
		if err != nil {
			err = errors.New("le dossier n'arrive pas à se créer")
			return "", err
		}
	}

	// création du fichier de destination
	imgPath = destination + strNomFichier
	dst, err := os.Create(imgPath)
	if err != nil {
		err = errors.New("le fichier n' a pas les droits pour être créé")
		return "", err
	}

	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		errors.New("le fichier n' a pas les droits pour être copié")
		return "", err
	}

	// envoie du path de l'image avec son nom
	return imgPath, nil
}

// cropImage permet de redimensionnement d'une image
// @todo il faut permettre le crop de jpg, png et gif
// pour le moment la fonction permet de réaliser des crop uniquemen pour des jpeg
func cropImage(imageUrl string, toW uint, toH uint) error {

	// open file
	rfile, err := os.Open(imageUrl)
	defer rfile.Close()
	if err != nil {
		return errors.New("le fichier n' a pas les droits pour être ouvert")
	}

	// decode jpeg into image.Image
	rimg, err := jpeg.Decode(rfile)
	if err != nil {
		return errors.New("l'image n'a pas put être décodé")
	}
	//rfile.Close()

	image := resize.Thumbnail(toW, toH, rimg, resize.Bilinear)

	out, err := os.Create(imageUrl)
	defer out.Close()
	if err != nil {
		return errors.New("l'image n'a pas put être créé")
	}

	// write new image to file
	jpeg.Encode(out, image, nil)

	return nil
}
