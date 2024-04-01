package cmd

import (
	"fmt"
	"log"

	templates "github.com/FelipeMCassiano/gorvus/gorvus/template"
	"github.com/spf13/cobra"
)

var createDockerfileCmd = &cobra.Command{
	Use:   "create Dockerfile [language]",
	Short: "Create Dockerfile based on input language and project name",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		language := args[1]
		createDockerfile(projectName, language)
		fmt.Println("Dockerfile created succesfully")
	},
}

func createDockerfile(projectName string, language string) {
	if language == "go" {
		if err := templates.GoDockerfile(projectName); err != nil {
			log.Fatalf("error: %s", err.Error())
			return
		}
	}
}
