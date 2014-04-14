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

type PageHome struct {
	Images [100]UserImage
	PageWeb
	Didi string
}
