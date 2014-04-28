package app

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
	RenderHtml    bool
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
