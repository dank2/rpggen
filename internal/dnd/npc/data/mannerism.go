package data

import (
	"math/rand"
	"time"
)

var Mannerism = map[int]string{
	1:  "Prone to singing, whistling, or humming quietly",
	2:  "Speaks in rhyme or some other peculiar way",
	3:  "Particularly low or high voice",
	4:  "Slurs words, lisps, or stutters",
	5:  "Enunciates overly clearly",
	6:  "Speaks loudly",
	7:  "Whispers",
	8:  "Uses flowery speech or long words",
	9:  "Frequently uses the wrong word",
	10: "Uses colorful oaths and exclamations",
	11: "Makes constant jokes or puns",
	12: "Prone to predictions of doom",
	13: "Fidgets",
	14: "Squints",
	15: "Stares into the distance",
	16: "Chews something",
	17: "Paces",
	18: "Taps fingers",
	19: "Bites fingernails",
	20: "Twirls hair or tugs beard",
}

func GetMannerism() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(Mannerism) + 1)

	return Mannerism[index]
}
