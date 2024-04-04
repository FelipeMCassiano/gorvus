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
	rootCmd.AddCommand(CreateComposeCommand())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
