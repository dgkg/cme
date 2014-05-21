package app

import (
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/kennygrant/sanitize"
	//"log"
	"net/http"
	. "strconv"
	//"time"
	"encoding/json"
)

type PageUserProject struct {
	User
	PagesList []Paginate
	PageWeb
}

// Public function
// permet de récupérer une liste de projets
func UserProjectAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	//page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCat(3)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}
		fmt.Fprint(w, string(b))
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
func UserProjectMoreAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if page == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		ups := up.getListPaged(page, 3)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}
		fmt.Fprint(w, string(b))
	}
}

// Public function
// permet de récupérer plus de projets
// avec de la pagination
// et avec une catégorie sélectionnée
func UserProjectMoreInCatAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	page, _ := Atoi(vars["page"])
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 || page == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCatPaged(page, 3)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}
		fmt.Fprint(w, string(b))
	}
}
