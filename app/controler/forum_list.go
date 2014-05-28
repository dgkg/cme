package controler

import (
	"github.com/kennygrant/sanitize"
	"log"
	"math"
	. "strconv"
	//"regexp"
	. "github.com/konginteractive/cme/app/model"
)

type PageForumList struct {
	Categories []ForumCategory
	PagesList  []Paginate
	Forums     []Forum
	PageWeb
}

// fonction pour permettre de créer une page
func CreatePageForm() *PageForum {
	return new(PageForum)
}

// fonction public
// permet d'afficher la liste des questions du forum
func (pf *PageForumList) View() {

	log.Println("ForumView appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pf.Title = "Forum"
	pf.MainClass = "forum"
	pf.Forums = f.GetList()
	pf.Categories = f.GetAllCategories(0)
	pf.PagesList = pf.createPaginate()
	pf.RenderHtml = true
	pf.injectDataToDisplay(pf.Forums)

}

// fonction public
// permet d'afficher la liste des questions du forum
// avec la fonction de pagination
func (pf *PageForumList) ViewPaged(page string) {

	log.Println("ForumViewPaged appelé : " + page)

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	//pagePosition, _ := ParseInt(page, 0, 64)
	pagePosition, _ := Atoi(page)

	pf.Title = "Forum"
	pf.MainClass = "forum"
	pf.Forums = f.GetListPaged(pagePosition)
	pf.Categories = f.GetAllCategories(0)
	pf.PagesList = pf.createPaginate()
	pf.RenderHtml = true
	pf.injectDataToDisplay(pf.Forums)

}

// fonction public
// permet d'afficher la liste des questions du forum avec la catégorie correspondante
func (pf *PageForumList) ViewCategory(cat string) {

	log.Println("ForumView appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	// récupère l'id de la catégorie
	idCat := f.GetIdFromCatName(cat)

	pf.Title = "Forum " + cat
	pf.MainClass = "forum"
	pf.Forums = f.GetListFromCat(idCat)
	pf.Categories = f.GetAllCategories(idCat)
	pf.PagesList = pf.createPaginateFromIdCat(idCat)
	pf.RenderHtml = true
	pf.injectDataToDisplay(pf.Forums)

}

// fonction public
// permet d'afficher une liste de questions en fonction de
// la catégorie sélectionnée
// et de la page en cours sélectionnée
func (pf *PageForumList) ViewCategoryPaged(cat string, page string) {

	log.Println("ForumView appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pagePosition, _ := Atoi(page)
	// récupère l'id de la catégorie
	idCat := f.GetIdFromCatName(cat)

	pf.Title = "Forum " + cat
	pf.MainClass = "forum"
	pf.Forums = f.GetListFromCatPaged(idCat, pagePosition)
	pf.Categories = f.GetAllCategories(idCat)
	pf.PagesList = pf.createPaginateFromIdCat(idCat)
	pf.RenderHtml = true
	pf.injectDataToDisplay(pf.Forums)

}

// fonction public
// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les titres
// et va chercher dans le texte du titre
func (pf *PageForumList) ViewSearch(q string) {

	log.Println("ForumViewSearch appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pf.Title = "Forum Rechercher"
	pf.MainClass = "forum"
	pf.Forums = f.Search(q)
	pf.Categories = f.GetAllCategories(0)
	pf.RenderHtml = true
	if len(pf.Forums) == 0 {
		pf.SearchText = q
	}
	pf.injectDataToDisplay(pf.Forums)

}

// fonction privée
// Permet de retrouver le nombre de réponses pour chaque post
// Permet aussi de réduire la description du texte de desc à 250 caractères
func (pf PageForumList) injectDataToDisplay(forums []Forum) []Forum {

	lenForum := len(forums)

	for i := 0; i < lenForum; i++ {
		id := forums[i].Id
		// permet de réaliser des extraits si le texte est trop long
		extrait := sanitize.HTML(forums[i].Text)
		if len(extrait) > 250 {
			extrait = extrait[0:250]
		}
		forums[i].Text = "<p>" + extrait + "</p>"
		// permet de compter ne nombres de réponses
		forums[i].PostNumb = forums[i].CountPost(id)
		// permet de créer une url du lien
		forums[i].Url = "/forum/post/" + Itoa(int(forums[i].Id))
	}
	return forums
}

// fonction privée
// fonction pour créer la pagination
func (pf PageForumList) createPaginate() []Paginate {

	var f Forum
	elTotal := f.Count()

	nb := elTotal / maxElementsInPage
	mf := int(math.Ceil(float64(nb)))
	mf++ // permet de réaliser la correction du nombre de pages à l'affichage
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
func (pf PageForumList) createPaginateFromIdCat(id int64) []Paginate {
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
func (p PageForumList) IsHtmlRender() bool {
	return p.RenderHtml
}

// permet d'injecter des donnés de cession dans l'utilisateur loggé
func (p *PageForumList) SetSessionData(u User) (v bool) {
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
