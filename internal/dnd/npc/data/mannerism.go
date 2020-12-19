package data

import (
	"math/rand"
	"time"
)

var Mannerism = []string{
	"Prone to singing, whistling, or humming quietly",
	"Speaks in rhyme or some other peculiar way",
	"Particularly low or high voice",
	"Slurs words, lisps, or stutters",
	"Enunciates overly clearly",
	"Speaks loudly",
	"Whispers",
	"Uses flowery speech or long words",
	"Frequently uses the wrong word",
	"Uses colorful oaths and exclamations",
	"Makes constant jokes or puns",
	"Prone to predictions of doom",
	"Fidgets",
	"Squints",
	"Stares into the distance",
	"Chews something",
	"Paces",
	"Taps fingers",
	"Bites fingernails",
	"Twirls hair or tugs beard",
}

func GetMannerism() string {
	rand.Seed(time.Now().UnixNano())
	return Mannerism[rand.Intn(len(Mannerism))]
}
