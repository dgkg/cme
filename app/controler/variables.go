package controler

// template utilisé par les pages affichés
// par défaut le template est form
var Templ = "index"

// donne le nombre d'éléments max à afficher par page
var maxElementsInPage = 3

// permet d'avoir un format de date à l'affichage
const dateLayout = "2 Jan 2006 à 3:04pm"

// fichier de destination des projets des élèves
/*
var T = time.Now()
var Y = T.Year()
var M = T.Month()
var YS = strconv.Itoa(Y)
var MS = M.String()
*/

var URL_PROJECT_IMAGES = "./img/projets/"

// max upload size file
const MAX_SIZE_FILE_UPLOAD = 100000
