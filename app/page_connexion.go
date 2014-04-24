package controler

import (
	"log"
)

func (pc PageConnexion) View() Page {

	log.Println("Connexion appelé")

	var p M.PageConn

	p.Title = "Connexion"
	p.MainClass = "connexion"

	return p
}
