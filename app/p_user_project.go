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

// Public function
// permet d'ajouter un commenaire sur une fonction
func UserProjectAjaxHandler(w http.ResponseWriter, r *http.Request) {
	// récupération de la variable id
	vars := mux.Vars(r)
	idCat, _ := ParseInt(vars["id"], 0, 64)
	// Validation des données
	// Si une des variables est vide, la func retourne un "error"
	// ce qui fait afficher un message d'erreur
	if idCat == 0 {
		// envoie un message d'erreur
		fmt.Fprint(w, "error")
	} else {

		var up UserProject
		up.UserProjectCategoryId = idCat
		ups := up.getByIdCat(6)
		// encodage en json des projets
		b, err := json.Marshal(ups)

		if err != nil {
			// envoie un message d'erreur
			fmt.Fprint(w, "error")
		}

		fmt.Fprint(w, string(b))
	}
}
