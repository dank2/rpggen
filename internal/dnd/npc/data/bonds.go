package data

import (
	"fmt"
	"math/rand"
)

var Bonds = []string{
	"Dedicated to fulfilling a personal life goal",
	"Protective of close family members",
	"Protective of colleagues or compatriots",
	"Loyal to a benefactor, patron, or employer",
	"Captivated by a romantic interest",
	"Drawn to a special place",
	"Protective of a sentimental keepsake",
	"Protective of a valuable possession",
	"Out for revenge",
}

func GetBond() string {
	length := len(Bonds)
	index := random.Intn(length + 1)
	if index == length {
		return fmt.Sprintf("%s, %s", Bonds[rand.Intn(length)], Bonds[rand.Intn(length)])
	}

	return Bonds[index]
}
