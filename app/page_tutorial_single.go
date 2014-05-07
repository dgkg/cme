package app

import (
	"log"
	. "strconv"
)

type PageTutorialSingle struct {
	Tutorial Tutorial
	PageWeb
}

// fonction public
// permet d'afficher la liste des tutoriels
func (pts *PageTutorialSingle) View(idTuto string) {

	log.Println("ViewSingle appelé")
	log.Println("Id appelé : " + Itoa(int(pts.Tutorial.Id)))

	// surcharge de la variable d'affichage
	Templ = "tutorial_single"
	pts.Title = "Tutoriel"
	pts.MainClass = "tutoriels"

	pts.Tutorial.Id, _ = ParseInt(idTuto, 0, 64)
	pts.Tutorial.getData()
	pts.Tutorial.Posts = pts.Tutorial.getPost()
	pts.RenderHtml = true
	//pts.injectDataToDisplay()

	log.Println("le titre du post est : " + pts.Tutorial.Title)

}

/*
*
*
*
*
*
*
*
*
*
*
*
*
*
 */
