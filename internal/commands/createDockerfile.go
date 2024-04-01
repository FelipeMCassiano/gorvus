package commands

import (
	"fmt"
	"log"

	"github.com/FelipeMCassiano/gorvus/internal/templates"
	"github.com/spf13/cobra"
)

var createDockerfileCmd = &cobra.Command{
	Use:   "createDockerfile <projectName> <templateLanguage>",
	Short: "Create Dockerfile based on input language and project name",
	Args:  cobra.ExactArgs(2),
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
