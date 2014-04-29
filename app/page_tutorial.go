package app

import (
	"log"
	"math"
	"net/http"
	. "strconv"
)

type PageTutorial struct {
	Categories []TutorialCategory
	PagesList  []Paginate
	Tutorials  []Tutorial
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageTutorial() *PageTutorial {
	return new(PageTutorial)
}

// fonction public
// permet d'afficher la liste des tutoriels
func (pt *PageTutorial) View() {

	log.Println("ForumView appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pt.Title = "Tutoriel"
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.getList()
	pt.Categories = t.getAllCategories()
	pt.PagesList = pt.createPaginate()
	pt.RenderHtml = false
	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher la liste des tutoriels
// avec la fonction de pagination
func (pt *PageTutorial) ViewPaged(page string) {

	log.Println("ForumViewPaged appelé : " + page)

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pagePosition, _ := ParseInt(page, 0, 64)

	pt.Title = "Tutoriel"
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.getListPaged(pagePosition)
	pt.Categories = t.getAllCategories()
	pt.PagesList = pt.createPaginate()
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher la liste des tutoriels avec la catégorie correspondante
func (pt *PageTutorial) ViewCategory(cat string) {

	log.Println("ForumView appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	// récupère l'id de la catégorie
	idCat := t.getIdFromCatName(cat)

	pt.Title = "Tutoriel " + cat
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.getListFromCat(idCat)
	pt.Categories = t.getAllCategories()
	pt.PagesList = pt.createPaginateFromIdCat(idCat)
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher une liste tutoriels en fonction de
// la catégorie sélectionnée
// et de la page en cours sélectionnée
func (pt *PageTutorial) ViewCategoryPaged(cat string, page string) {

	log.Println("ForumView appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pagePosition, _ := ParseInt(page, 0, 64)
	// récupère l'id de la catégorie
	idCat := t.getIdFromCatName(cat)

	pt.Title = "Tutoriel " + cat
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.getListFromCatPaged(idCat, pagePosition)
	pt.Categories = t.getAllCategories()
	pt.PagesList = pt.createPaginateFromIdCat(idCat)
	pt.RenderHtml = true
	pt.injectDataToDisplay(pt.Tutorials)
}

// fonction public
// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les titres
// et va chercher dans le texte du titre
func (pt *PageTutorial) ViewSearch(q string) {

	log.Println("ForumViewSearch appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial"

	pt.Title = "Tutoriel Rechercher"
	pt.MainClass = "tutoriels"
	pt.Tutorials = t.search(q)
	pt.Categories = t.getAllCategories()
	pt.RenderHtml = true
	if len(pt.Tutorials) == 0 {
		pt.SearchText = q
	}

	pt.injectDataToDisplay(pt.Tutorials)

}

// fonction public
// permet d'afficher le formulaire de création d'un tutoriel
func (pt *PageTutorial) ViewAdd() {

	log.Println("ViewAdd appelé")

	var t Tutorial

	// surcharge de la variable d'affichage
	Templ = "tutorial_add"

	pt.Title = "Titre du sujet"
	pt.MainClass = "nouveaututo"
	pt.Tutorials = make([]Tutorial, 1)
	pt.Categories = t.getAllCategories()
	pt.RenderHtml = true
}

// fonction public
// permet de valider les éléments du formulaire
func (pt PageTutorial) ValidateForm(r *http.Request) (t Tutorial, v bool) {

	log.Println("ValidateForm appelé")

	t.Title = r.PostFormValue("post-nom")
	t.Text = r.PostFormValue("post-contenu")
	t.TutorialCategoryId, _ = ParseInt(r.PostFormValue("post-cat"), 0, 64)
	t.IsOnline = 1 // permet de mettre en ligne automatiquement la quesion

	// on initialise v à true
	v = true
	// on vérifi les champs qui sont présents
	if r.PostFormValue("post-nom") == "" {
		v = false
	}

	if r.PostFormValue("post-contenu") == "" {
		v = false
	}

	return
}

// fonction privée
// Permet de retrouver le nombre de réponses pour chaque post
// Permet aussi de réduire la description du texte de desc à 250 caractères
func (pt PageTutorial) injectDataToDisplay(tutorials []Tutorial) []Tutorial {

	lenTuto := len(tutorials)

	for i := 0; i < lenTuto; i++ {
		id := tutorials[i].Id
		// permet de réaliser des extraits si le texte est trop long
		if len(tutorials[i].Text) > 250 {
			text := tutorials[i].Text[0:250]
			tutorials[i].Text = text
		}
		tutorials[i].PostNumb = tutorials[i].countPost(id)
	}

	return tutorials
}

// fonction privée
// fonction pour créer la pagination
func (pt PageTutorial) createPaginate() []Paginate {

	var t Tutorial

	elTotal := t.count()

	nb := elTotal / maxElementsInPage
	mf := int(math.Floor(float64(nb)))
	p := make([]Paginate, nb)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
}

// fonction privée
// fonction pour créer la pagination à partir d'une catégorie sélectionnée
func (pt PageTutorial) createPaginateFromIdCat(id int64) []Paginate {
	var f Forum
	elTotal := f.countFromIdCat(id)

	nb := elTotal / maxElementsInPage
	mf := int(math.Floor(float64(nb)))
	p := make([]Paginate, nb)

	for i := 0; i < mf; i++ {
		t := Itoa(i + 1)
		p[i].Title = t
		p[i].Url = "?p=" + t
	}
	return p
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p PageTutorial) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageTutorial) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
}
