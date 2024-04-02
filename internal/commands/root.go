package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gorvus",
	Short:   "",
	Long:    ``,
	Version: "1.0.0",
}

func Execute() {
	rootCmd.AddCommand(createDockerfile())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
