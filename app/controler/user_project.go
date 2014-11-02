package controler

import (
	. "github.com/konginteractive/cme/app/model"
)

type PageUserProject struct {
	User
	PagesList []Paginate
	PageWeb
}
