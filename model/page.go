package model

type Page interface {
}

type PageWeb struct {
	Title         string
	Keywords      []Keyword
	VarSessServer string
	VarSessCookie string
	VarFlash      string
	MainClass     string
	SearchText    string // permet de récupérer dans une vue le(s) mot(s) clefs recherchés
	Logge         bool
	Menu
}

type Keyword struct {
	Title string
}

type Menu struct {
}

type Paginate struct {
	Title string
	Url   string
}

type PageHome struct {
	Images []UserImage
	PageWeb
}

type PageForum struct {
	Categories []ForumCategory
	PagesList  []Paginate
	Forums     []Forum
	PageWeb
}

type PageNews struct {
	PagesList []Paginate
	News      []NewsViewHelper
	PageWeb
}

type PageTutoriels struct {
	PagesList []Paginate
	Tutoriels []TutorialsViewHelper
	PageWeb
}

type PageUser struct {
	PagesList []Paginate
	Users     []UsersViewHelper
	PageWeb
}
