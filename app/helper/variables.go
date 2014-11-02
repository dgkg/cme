package helper

import (
	"strconv"
	"time"
)

// fichier de destination des projets des élèves
var T = time.Now()
var Y = T.Year()
var M = T.Month()
var YS = strconv.Itoa(Y)
var MS = M.String()

// max upload size file
const MAX_SIZE_FILE_UPLOAD = 100000
