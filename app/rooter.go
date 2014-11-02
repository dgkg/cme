package app

import (
	. "github.com/konginteractive/cme/app/controler"
	"log"
	"net/http"
)

func DispatcherHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Referer())
}
