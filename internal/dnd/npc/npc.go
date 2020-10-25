package npc

import (
	"math/rand"
	"time"

	data "github.com/dank2/rpggen/internal/dnd/npc/data"
)

// Npc is the Base type NPC
type Npc struct {
	Name        string `json:name,omitempty`
	Race        string `json:race,omitempty`
	Gender      string `json:gender,omitempty`
	Personality string `json:personality,omitempty`
	Appearance  string `json:appearance,omitempty`
	Dream       string `json:dream,omitempty`
	Flaw        string `json:flaw,omitempty`
}

// GenerateNpc is theMain func to generate an NPC
func GenerateNpc() (*Npc, error) {

	gender := data.GetGender()
	name := data.GetName(gender)

	rand.Seed(time.Now().UnixNano())
	return &Npc{
		Name:   name,
		Gender: gender,
	}, nil
}
