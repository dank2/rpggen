package data

import (
	"math/rand"
	"time"
)

var Appearances = map[int]string{
	1:  "Distinctive jewelry: earrings, necklace, circlet, bracelets",
	2:  "Piercings",
	3:  "Flamboyant or outlandish clothes",
	4:  "Formal, clean clothes",
	5:  "Ragged, dirty clothes",
	6:  "Pronounced scar",
	7:  "Missing teeth",
	8:  "Missing fingers",
	9:  "Unusual eye color (or two different colors)",
	10: "Tattoos",
	11: "Birthmark",
	12: "Unusual skin color",
	13: "Bald",
	14: "Braided beard or hair",
	15: "Unusual hair color",
	16: "Nervous eye twitch",
	17: "Distinctive nose",
	18: "Distinctive posture (crooked or rigid)",
	19: "Exceptionally beautiful",
	20: "Exceptionally ugly",
}

func GetAppearance() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(Appearances) + 1)

	return Appearances[index]
}
