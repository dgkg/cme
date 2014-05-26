package app

import (
	"strconv"
	"time"
)

// template utilisé par les pages affichés
// par défaut le template est form
var Templ = "index"

// donne le nombre d'éléments max à afficher par page
var maxElementsInPage = 3

// connexion à la base de donnée
var db = connectToDatabase()

// permet d'avoir un format de date à l'affichage
const dateLayout = "2 Jan 2006 à 3:04pm"

// fichier de destination des projets des élèves
var T = time.Now()
var Y = T.Year()
var M = T.Month()
var YS = strconv.Itoa(Y)
var MS = M.String()

var URL_PROJECT_IMAGES = "./img/projets/"
