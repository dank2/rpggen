package npc

import (
	"fmt"
	"math/rand"
	"time"

	data "github.com/dank2/rpggen/internal/dnd/npc/data"
)

// Npc is the Base type NPC
type Npc struct {
	Name       string `yaml:"Name,omitempty"`
	Race       string `yaml:"Race,omitempty"`
	Mannerism  string `yaml:"Mannerism,omitempty"`
	Appearance string `yaml:"Appearance,omitempty"`
	Bond       string `yaml:"Bond,omitempty"`
	Flaw       string `yaml:"Flaw,omitempty"`
}

// GenerateNpc is theMain func to generate an NPC
func GenerateNpc() (*Npc, error) {

	gender := data.GetGender()
	race := data.GetRace()
	name := data.GetName(gender, race)
	appearance := data.GetAppearance()
	mannerism := data.GetMannerism()
	bond := data.GetBond()
	flaw := data.GetFlaw()

	rand.Seed(time.Now().UnixNano())
	return &Npc{
		Name:       name,
		Race:       race,
		Appearance: fmt.Sprintf("%s, %s", gender, appearance),
		Mannerism:  mannerism,
		Bond:       bond,
		Flaw:       flaw,
	}, nil
}
