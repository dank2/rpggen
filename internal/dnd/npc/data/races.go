package data

import (
	"math/rand"
	"time"
)

var Races = map[int]string{
	1: "human",
	2: "dwarf",
	3: "gnome",
	4: "halfling",
	5: "elf",
	6: "half-elf",
	7: "dragonborn",
}

func GetRace() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(Races) + 1)

	return Races[index]
}
