package model

type Page interface {
}

type PageWeb struct {
	Title         string
	Keywords      []Keyword
	VarSessServer string
	VarSessCookie string
	VarFlash      string
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
	Images [100]UserImage
	PageWeb
}

type PageForum struct {
	PagesList []Paginate
	Forums    []ForumViewHelper
	PageWeb
}
