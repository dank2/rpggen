package cmd

import (
	"io"

	"github.com/dank2/rpggen/rpggen/cmd/dnd"
	"github.com/spf13/cobra"
)

func AddDndCommand(out, err io.Writer) *cobra.Command {
	dndCommand := &cobra.Command{
		Use:   "dnd",
		Short: "The dnd subpackage.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	dndCommand.AddCommand(dnd.AddNpcCommand(out))

	return dndCommand
}
