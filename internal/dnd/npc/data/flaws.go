package data

import (
	"math/rand"
	"time"
)

var Flaws = []string{
	"Forbidden love or susceptibility to romance",
	"Enjoys decadent pleasures",
	"Arrogance",
	"Envies another creature’s possessions or station",
	"Overpowering greed",
	"Prone to rage",
	"Has a powerful enemy",
	"Specific phobia",
	"Shameful or scandalous history",
	"Secret crime or misdeed",
	"Possession of forbidden lore",
	"Foolhardy bravery",
}

func GetFlaw() string {
	rand.Seed(time.Now().UnixNano())
	return Flaws[rand.Intn(len(Flaws))]
}
