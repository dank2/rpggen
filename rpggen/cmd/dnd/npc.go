package dnd

import (
	"fmt"
	"io"

	"github.com/dank2/rpggen/internal/dnd/npc"
	"github.com/dank2/rpggen/internal/utils"
	"github.com/spf13/cobra"
)

func AddNpcCommand(out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "npc",
		Short: "generate a dnd npc",
		RunE: func(cmd *cobra.Command, args []string) error {
			npc, err := npc.GenerateNpc()
			if err != nil {
				return fmt.Errorf("unable to generate npc: %s\n", err)
			}

			if formatted, err := utils.PrettyFormat(npc); err == nil {
				fmt.Fprintf(out, "Generated NPC:\n%s", formatted)
			} else if err != nil {
				return fmt.Errorf("unable to generate npc: %s", err)
			}

			return nil
		},
	}
}
