package app

import ()

/*
// Helper de vue
type NewsViewHelper struct {
	CategoryTitle string
	News
}
*/

type PageNews struct {
	PagesList []Paginate
	News      []NewsViewHelper
	PageWeb
}

func (pn PageNews) View() Page {

	// surcharge de la variable d'affichage
	Templ = "news"

	db, _ := gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/cme_test?charset=utf8&parseTime=True")
	//cme_test
	db.SingularTable(true)

	log.Println("Forum appelé")

	pn.Title = "News"
	pn.MainClass = "actualites"

	// pagination
	pn.PagesList = make([]Paginate, 5)

	pn.PagesList[0].Title = "1"
	pn.PagesList[0].Url = "/forum/page/1"

	pn.PagesList[1].Title = "2"
	pn.PagesList[1].Url = "/forum/page/2"

	pn.PagesList[2].Title = "3"
	pn.PagesList[2].Url = "/forum/page/3"

	pn.PagesList[3].Title = "4"
	pn.PagesList[3].Url = "/forum/page/4"

	pn.PagesList[4].Title = "5"
	pn.PagesList[4].Url = "/forum/page/5"

	return pn
}
