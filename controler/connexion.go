package controler

import (
	M "github.com/konginteractive/cme/model"
	"log"
)

var ConnexionTempl = "connexion"

func ConnexionView() M.Page {

	log.Println("Connexion appelé")

	var p M.PageConn

	p.Title = "Connexion"
	p.MainClass = "connexion"

	return p
}
