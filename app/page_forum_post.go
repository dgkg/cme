package app

import (
	"log"
)

type PageForumPost struct {
	//PagesList []Paginate
	//News      []NewsViewHelper
	PageWeb
}

func (pfp PageForumPost) View() Page {

	log.Println("Forum Post appelé")

	// surcharge de la variable d'affichage
	Templ = "forum_post"

	pfp.Title = "Titre du post! Woohoo!"
	pfp.MainClass = "forum_post"

	return pfp
}
