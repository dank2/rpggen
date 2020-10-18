package dnd

import (
	"math/rand"
	"time"
)

func GetGender() string {
	rand.Seed(time.Now().UnixNano())
	genderNumber := rand.Intn(1)

	if genderNumber == 1 {
		return "male"
	} else {
		return "female"
	}
}
