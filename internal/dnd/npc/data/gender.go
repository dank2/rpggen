package data

import (
	"math/rand"
	"time"
)

func GetGender() string {
	rand.Seed(time.Now().UnixNano())
	genderNumber := rand.Intn(2)

	if genderNumber == 1 {
		return "male"
	} else {
		return "female"
	}
}
