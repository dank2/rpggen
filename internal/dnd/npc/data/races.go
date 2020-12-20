package data

import (
	"sort"
)

type RacesType []RaceDist

var Races RacesType = []RaceDist{
	{"Human", 0.5},
	{"Dwarf", 0.1},
	{"Halfling", 0.1},
	{"Gnome", 0.06},
	{"Half-elf", 0.06},
	{"Elf", 0.06},
	{"Dragonborn", 0.06},
	{"Half-orc", 0.06},
}

type RaceDist struct {
	Race         string
	Distribution float32
}

func GetRace() string {
	Races.Sort()

	pointer := random.Float32()
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
