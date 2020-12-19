package npc

import (
	"fmt"
	"math/rand"
	"time"

	data "github.com/dank2/rpggen/internal/dnd/npc/data"
)

// Npc is the Base type NPC
type Npc struct {
	Name       string `json:name,omitempty`
	Race       string `json:race,omitempty`
	Mannerism  string `json:mannerism,omitempty`
	Appearance string `json:appearance,omitempty`
	Bond       string `json:bond,omitempty`
	Flaw       string `json:flaw,omitempty`
}

// GenerateNpc is theMain func to generate an NPC
func GenerateNpc() (*Npc, error) {

	gender := data.GetGender()
	race := data.GetRace()
	name := data.GetName(gender)
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
