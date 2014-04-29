package app

type Page interface {
	IsHtmlRender() bool
	SetSessionData(u User) (v bool)
}

type PageWeb struct {
	Title         string
	Keywords      []Keyword
	SessNameUser  string
	SessIsLogged  bool
	VarSessServer string
	VarSessCookie string
	VarFlash      string
	MainClass     string
	SearchText    string // permet de récupérer dans une vue le(s) mot(s) clefs recherchés
	Logge         bool
	RenderHtml    bool
	Menu
}

// fonction permettant de savoir si le rendu passe par l'html ou non
// permet de faire fonctionner avec l'interface de type Page
func (p *PageWeb) IsHtmlRender() bool {
	return p.RenderHtml
}

func (p *PageWeb) SetSessionData(u User) (v bool) {
	if u.Id != 0 && u.FirstName != "" {
		p.SessIsLogged = true
		p.SessNameUser = u.FirstName
		v = true
	}
	return
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
