package controler

import (
	M "github.com/konginteractive/cme/model"
	"log"
)

var HomeTempl = "home"

func HomeView() M.Page {

	log.Println("Home appelé")

	p := new(M.PageHome)
	p.Title = "Coucouc de la forêt"

	p.MainClass = "accueil"

	/*
		p.Images[0].Url = "http://placekitten.com/100/600"
		p.Images[1].Url = "http://placekitten.com/200/50"
	*/
	// Works !!
	//p.Images = []M.UserImage{{Url: "http://placekitten.com/600/600"}, {Url: "http://placekitten.com/600/300"}}
	//connectToDatabase()
	return p
}
