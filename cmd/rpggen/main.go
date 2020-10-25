package main

import (
	"encoding/json"
	"fmt"

	"github.com/dank2/rpggen/internal/dnd/npc"
)

func main() {
	generatedNpc, _ := npc.GenerateNpc()

	npcJSON, _ := json.MarshalIndent(generatedNpc, "", "  ")

	fmt.Printf(string(npcJSON))
}
