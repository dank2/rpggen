package dnd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddNpcCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "npc",
		Short: "generate a dnd npc",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("generating a dnd npc...")
			return nil
		},
	}

}
