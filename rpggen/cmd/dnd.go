package cmd

import (
	"github.com/dank2/rpggen/rpggen/cmd/dnd"
	"github.com/spf13/cobra"
)

var dndCommand = &cobra.Command{
	Use:   "dnd",
	Short: "The dnd subpackage.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	dndCommand.AddCommand(dnd.AddNpcCommand())

	rootCmd.AddCommand(dndCommand)
}
