package data

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
}
