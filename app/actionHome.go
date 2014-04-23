package app

import (
	"log"
)

func (ph PageHome) View() Page {

	log.Println("Home appelé")

	// surcharge de la variable d'affichage
	Templ = "index"

	ph.Title = "Coucouc de la forêt"

	ph.MainClass = "accueil"

	/*
		p.Images[0].Url = "http://placekitten.com/100/600"
		p.Images[1].Url = "http://placekitten.com/200/50"
	*/
	// Works !!
	//p.Images = []M.UserImage{{Url: "http://placekitten.com/600/600"}, {Url: "http://placekitten.com/600/300"}}
	//connectToDatabase()
	return ph

}
