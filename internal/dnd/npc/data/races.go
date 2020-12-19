package data

import (
	"math/rand"
	"time"
)

var Races = []string{
	"Human",
	"Dwarf",
	"Gnome",
	"Halfling",
	"Elf",
	"Half-elf",
	"Dragonborn",
	"Half-orc",
}

func GetRace() string {
	rand.Seed(time.Now().UnixNano())
	return Races[rand.Intn(len(Races))]
}
