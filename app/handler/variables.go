package handler

// permet d'avoir un format de date à l'affichage
const dateLayout = "2 Jan 2006 à 3:04pm"

// fichier de destination des projets des élèves
var T = time.Now()
var Y = T.Year()
var M = T.Month()
var YS = strconv.Itoa(Y)
var MS = M.String()
