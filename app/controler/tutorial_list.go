package controler

import (
	"github.com/kennygrant/sanitize"
	. "github.com/konginteractive/cme/app/model"
	"log"
	"math"
	. "strconv"
)

type PageTutorialList struct {
	Categories []TutorialCategory
	PagesList  []Paginate
	Tutorials  []Tutorial
	PageWeb
}

// fonction public
// permet d'afficher la liste des tutoriels
func (pt *PageTutorialList) View() {

	log.Println("ForumView appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pt.Title = "Tutoriel"
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.GetList()
	pt.Categories = t.GetAllCategories(0)
	pt.PagesList = pt.createPaginate()
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher la liste des tutoriels
// avec la fonction de pagination
func (pt *PageTutorialList) ViewPaged(page string) {

	log.Println("ForumViewPaged appelé : " + page)

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pagePosition, _ := Atoi(page)

	pt.Title = "Tutoriel"
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.GetListPaged(pagePosition)
	pt.Categories = t.GetAllCategories(0)
	pt.PagesList = pt.createPaginate()
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher la liste des tutoriels avec la catégorie correspondante
func (pt *PageTutorialList) ViewCategory(cat string) {

	log.Println("ForumView appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	// récupère l'id de la catégorie
	idCat := t.GetIdFromCatName(cat)

	pt.Title = "Tutoriel " + cat
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.GetListFromCat(idCat)
	pt.Categories = t.GetAllCategories(idCat)
	pt.PagesList = pt.createPaginateFromIdCat(idCat)
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher une liste tutoriels en fonction de
// la catégorie sélectionnée
// et de la page en cours sélectionnée
func (pt *PageTutorialList) ViewCategoryPaged(cat string, page string) {

	log.Println("ForumView appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pagePosition, _ := Atoi(page)
	// récupère l'id de la catégorie
	idCat := t.GetIdFromCatName(cat)

	pt.Title = "Tutoriel " + cat
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.GetListFromCatPaged(idCat, pagePosition)
	pt.Categories = t.GetAllCategories(idCat)
	pt.PagesList = pt.createPaginateFromIdCat(idCat)
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)
}

// fonction public
// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les titres
// et va chercher dans le texte du titre
func (pt *PageTutorialList) ViewSearch(q string) {

	log.Println("ForumViewSearch appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pt.Title = "Tutoriel Rechercher"
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.Search(q)
	pt.Categories = t.GetAllCategories(0)
	pt.RenderHtml = true
	if len(pt.Tutorials) == 0 {
		pt.SearchText = q
	}

	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction privée
// Permet de retrouver le nombre de réponses pour chaque post
// Permet aussi de réduire la description du texte de desc à 250 caractères
func (pt PageTutorialList) injectDataToDisplay(tutorials []Tutorial) []Tutorial {

	lenTuto := len(tutorials)

	for i := 0; i < lenTuto; i++ {
		id := tutorials[i].Id
		// permet de réaliser des extraits si le texte est trop long
		extrait := sanitize.HTML(tutorials[i].Text)
		if len(extrait) > 250 {
			extrait = extrait[0:250]
		}
		tutorials[i].Text = "<p>" + extrait + "</p>"
		// permet de compter ne nombres de réponses
		tutorials[i].PostNumb = tutorials[i].CountPost(id)
		// permet de créer une url du lien
		tutorials[i].CategoryTitle = tutorials[i].GetCatTitle()
	}
	return tutorials
}

// fonction privée
// fonction pour créer la pagination
func (pt PageTutorialList) createPaginate() []Paginate {

	var t Tutorial

	elTotal := t.Count()

	nb := elTotal / maxElementsInPage
	mf := int(math.Ceil(float64(nb)))
	mf++ // permet de réaliser la correction du nombre de pages
	p := make([]Paginate, mf)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
}

// fonction privée
// fonction pour créer la pagination à partir d'une catégorie sélectionnée
func (pt PageTutorialList) createPaginateFromIdCat(id int64) []Paginate {
	var f Forum
	elTotal := f.CountFromIdCat(id)

	nb := elTotal / maxElementsInPage
	mf := int(math.Ceil(float64(nb)))
	mf++ // permet de réaliser la correction du nombre de pages
	p := make([]Paginate, mf)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageTutorialList) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageTutorialList) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIdUser = u.Id
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	} else {
		p.SessIdUser = 0
		p.SessIsLogged = false
		p.SessNameUser = ""
	}
	return
}
