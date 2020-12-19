package data

import (
	"math/rand"
	"time"
)

var Appearances = []string{
	"Distinctive jewelry: earrings, necklace, circlet, bracelets",
	"Piercings",
	"Flamboyant or outlandish clothes",
	"Formal, clean clothes",
	"Ragged, dirty clothes",
	"Pronounced scar",
	"Missing teeth",
	"Missing fingers",
	"Unusual eye color (or two different colors)",
	"Tattoos",
	"Birthmark",
	"Unusual skin color",
	"Bald",
	"Braided beard or hair",
	"Unusual hair color",
	"Nervous eye twitch",
	"Distinctive nose",
	"Distinctive posture (crooked or rigid)",
	"Exceptionally beautiful",
	"Exceptionally ugly",
}

func GetAppearance() string {
	rand.Seed(time.Now().UnixNano())
	return Appearances[rand.Intn(len(Appearances))]
}
