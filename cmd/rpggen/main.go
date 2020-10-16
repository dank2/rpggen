package main

import (
	"fmt"

	"github.com/dank2/rpggen/internal/npc"
)

func main() {
	fmt.Println("Hello world")

	generatedNpc, _ := npc.GenerateNpc()

	fmt.Printf("First NPC name: %s\n", generatedNpc.Name)
}
