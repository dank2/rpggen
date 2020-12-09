package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "rpggen",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, ars []string) {
			fmt.Println("ehehehe")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

}
