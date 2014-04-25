package app

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"net/http"
	. "strconv"
)

type PageForum struct {
	Categories []ForumCategory
	PagesList  []Paginate
	Forums     []Forum
	PageWeb
}

// fonction public
// permet d'afficher la liste des questions du forum
func (pf PageForum) View() Page {

	log.Println("ForumView appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pf.Title = "Forum"
	pf.MainClass = "forum"
	pf.Forums = f.getList()
	pf.Categories = f.getAllCategories()
	pf.PagesList = pf.createPaginate()

	pf.injectDataToDisplay(pf.Forums)

	return pf
}

// fonction public
// permet d'afficher la liste des questions du forum
// avec la fonction de pagination
func (pf PageForum) ViewPaged(page string) Page {

	log.Println("ForumViewPaged appelé : " + page)

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pagePosition, _ := ParseInt(page, 0, 64)

	pf.Title = "Forum"
	pf.MainClass = "forum"
	pf.Forums = f.getListPaged(pagePosition)
	pf.Categories = f.getAllCategories()
	pf.PagesList = pf.createPaginate()

	pf.injectDataToDisplay(pf.Forums)

	return pf
}

// fonction public
// permet d'afficher la liste des questions du forum avec la catégorie correspondante
func (pf PageForum) ViewCategory(cat string) Page {

	log.Println("ForumView appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	// récupère l'id de la catégorie
	idCat := f.getIdFromCatName(cat)

	pf.Title = "Forum " + cat
	pf.MainClass = "forum"
	pf.Forums = f.getListFromCat(idCat)
	pf.Categories = f.getAllCategories()
	pf.PagesList = pf.createPaginateFromIdCat(idCat)

	pf.injectDataToDisplay(pf.Forums)

	return pf
}

// fonction public
// permet d'afficher une liste de questions en fonction de
// la catégorie sélectionnée
// et de la page en cours sélectionnée
func (pf PageForum) ViewCategoryPaged(cat string, page string) Page {

	log.Println("ForumView appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pagePosition, _ := ParseInt(page, 0, 64)
	// récupère l'id de la catégorie
	idCat := f.getIdFromCatName(cat)

	pf.Title = "Forum " + cat
	pf.MainClass = "forum"
	pf.Forums = f.getListFromCatPaged(idCat, pagePosition)
	pf.Categories = f.getAllCategories()
	pf.PagesList = pf.createPaginateFromIdCat(idCat)

	pf.injectDataToDisplay(pf.Forums)

	return pf
}

// fonction public
// permet d'afficher la recherche à partir d'un mot clef
// va chercher dans les titres
// et va chercher dans le texte du titre
func (pf PageForum) ViewSearch(q string) Page {

	log.Println("ForumViewSearch appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum"

	pf.Title = "Forum Rechercher"
	pf.MainClass = "forum"
	pf.Forums = f.search(q)
	pf.Categories = f.getAllCategories()

	if len(pf.Forums) == 0 {
		pf.SearchText = q
	}

	pf.injectDataToDisplay(pf.Forums)

	return pf
}

// fonction public
// permet d'afficher le formulaire de création d'une question du formulaire
func (pf PageForum) ViewAdd() Page {

	log.Println("ViewAdd appelé")

	var f Forum

	// surcharge de la variable d'affichage
	Templ = "forum_add"

	pf.Title = "Créer un nouveau sujet"
	pf.MainClass = "nouveausujet"
	pf.Forums = make([]Forum, 1)
	pf.Categories = f.getAllCategories()

	return pf
}

// fonction public
// permet de valider les éléments du formulaire
func (pf PageForum) ValidateForm(r *http.Request) (f Forum, v bool) {

	log.Println("ValidateForm appelé")

	f.Title = r.PostFormValue("post-nom")
	f.Text = r.PostFormValue("post-contenu")
	f.ForumCategoryId, _ = ParseInt(r.PostFormValue("post-cat"), 0, 64)
	f.IsOnline = 1 // permet de mettre en ligne automatiquement la quesion

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
func (pf PageForum) injectDataToDisplay(forums []Forum) []Forum {

	lenForum := len(forums)

	for i := 0; i < lenForum; i++ {
		id := forums[i].Id
		// permet de réaliser des extraits si le texte est trop long
		if len(forums[i].Text) > 250 {
			text := forums[i].Text[0:250]
			forums[i].Text = text
		}
		// permet de compter ne nombres de réponses
		forums[i].PostNumb = forums[i].countPost(id)
		// permet de créer une url du lien
		forums[i].Url = "/forum/post/" + Itoa(int(forums[i].Id))
	}

	return forums
}

// fonction privée
// fonction pour créer la pagination
func (pf PageForum) createPaginate() []Paginate {
	var f Forum
	elTotal := f.count()

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
func (pf PageForum) createPaginateFromIdCat(id int64) []Paginate {
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
