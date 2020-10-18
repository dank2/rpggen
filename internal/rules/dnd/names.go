package dnd

import (
	"math/rand"
	"time"
)

var norseNamesMale = map[int]string{
	1:  "Agni",
	2:  "Alaric",
	3:  "Anvindr",
	4:  "Arvid",
	5:  "Asger",
	6:  "Asmund",
	7:  "Bjarte",
	8:  "Bjorg",
	9:  "Bjorn",
	10: "Brandr",
}

var norseNamesFemale = map[int]string{
	1:  "Alfhild",
	2:  "Arnbjorg",
	3:  "Ase",
	4:  "Aslog",
	5:  "Astrid",
	6:  "Auda",
	7:  "Audhid",
	8:  "Bergljot",
	9:  "Birghild",
	10: "Bodil",
}

// GetName returns a name from the table based on index
func GetName(gender string) string {
	rand.Seed(time.Now().UnixNano())

	var listToUse map[int]string
	if gender == "male" {
		listToUse = norseNamesMale
	} else {
		listToUse = norseNamesFemale
	}

	index := rand.Intn(len(listToUse) + 1)

	return listToUse[index]
}
