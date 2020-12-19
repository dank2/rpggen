package data

import (
	"math/rand"
	"time"
)

var norseNamesMale = []string{
	"Agni",
	"Alaric",
	"Anvindr",
	"Arvid",
	"Asger",
	"Asmund",
	"Bjarte",
	"Bjorg",
	"Bjorn",
	"Brandr",
}

var norseNamesFemale = []string{
	"Alfhild",
	"Arnbjorg",
	"Ase",
	"Aslog",
	"Astrid",
	"Auda",
	"Audhid",
	"Bergljot",
	"Birghild",
	"Bodil",
}

// GetName returns a name from the table based on index
func GetName(gender string) string {
	rand.Seed(time.Now().UnixNano())

	var listToUse []string
	if gender == "male" {
		listToUse = norseNamesMale
	} else {
		listToUse = norseNamesFemale
	}

	return listToUse[rand.Intn(len(listToUse))]
}
