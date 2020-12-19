package main

import (
	"fmt"
	"os"

	"github.com/dank2/rpggen/rpggen/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand(os.Stdout, os.Stderr)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
