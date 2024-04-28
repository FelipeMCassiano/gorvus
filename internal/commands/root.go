package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gorvus",
	Short: "gorvus is a CLI tool written in Go for generating Dockerfiles and docker-compose.yml files ",
	Long: `gorvus is a CLI tool written in Go that simplifies the process of generating Dockerfiles and docker-compose.yml files for your projects. 
    With gorvus, you can quickly scaffold Docker configurations without manual`,
	Version: "1.2.0",
}

func Execute() {
	rootCmd.AddCommand(generateDockerfile())
	rootCmd.AddCommand(CreateComposeCommand())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
