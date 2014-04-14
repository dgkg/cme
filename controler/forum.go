package controler

import (
	M "github.com/konginteractive/cme/model"
	"log"
)

var ForumTempl = "forum"

func ForumView() M.Page {

	log.Println("Forum appelé")

	p := new(M.PageHome)
	p.Title = "Coucouc de la forêt"

	p.Images[0].Url = "http://placekitten.com/100/600"
	p.Images[1].Url = "http://placekitten.com/200/50"

	// Works !!
	//p.Images = []M.UserImage{{Url: "http://placekitten.com/600/600"}, {Url: "http://placekitten.com/600/300"}}

	return p
}
