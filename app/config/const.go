package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// for production data put true
var Prod bool = false

var Data = &DbConf{}

type DbConf struct {
	DbType     []string
	DbUser     []string
	DbPass     []string
	DbProtocol []string
	DbName     []string
}

func init() {
	file, _ := os.Open("./db.json")
	decoder := json.NewDecoder(file)

	decoder.Decode(&Data)

	if Prod {
		fmt.Println(Data.DbPass[1])
	} else {
		fmt.Println(Data.DbPass[0])
	}

}
