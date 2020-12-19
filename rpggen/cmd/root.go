package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func NewRootCommand(out, err io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rpggen",
		Short: "rpggen is a tool to generate resources for ttrpgs",
		Long:  ``,
		RunE: func(cmd *cobra.Command, ars []string) error {
			return cmd.Help()
		},
	}

	rootCmd.AddCommand(AddDndCommand(out, err))
	return rootCmd
}
