package controler

import (
	M "github.com/ritoon/kong/cme_front/model"
	"log"
)

var HomeTempl = "home"

func HomeView() M.Page {

	log.Println("Home appelé")

	p := new(M.PageHome)
	p.Title = "Coucouc de la forêt"

	p.Images[0].Url = "http://placekitten.com/100/600"
	p.Images[1].Url = "http://placekitten.com/200/50"

	// Works !!
	//p.Images = []M.UserImage{{Url: "http://placekitten.com/600/600"}, {Url: "http://placekitten.com/600/300"}}

	return p
}
