package data

import (
	"math/rand"
	"sort"
	"time"
)

type RacesType []RaceDist

var Races RacesType = []RaceDist{
	{"Human", 0.8},
	{"Dwarf", 0.05},
	{"Halfling", 0.05},
	{"Gnome", 0.02},
	{"Half-elf", 0.02},
	{"Elf", 0.02},
	{"Dragonborn", 0.02},
	{"Half-orc", 0.02},
}

type RaceDist struct {
	Race         string
	Distribution float32
}

var r *rand.Rand

func GetRace() string {
	Races.Sort()

	pointer := r.Float32()
	var sum float32
	for _, v := range Races {
		sum += v.Distribution
		if pointer < sum {
			return v.Race
		}
	}

	return "Human"
}

func (r RacesType) GetValues() []string {
	result := make([]string, len(r))
	for i, v := range r {
		result[i] = v.Race
	}
	return result
}

func (r RacesType) Sort() {
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Distribution > r[j].Distribution
	})
}

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}
