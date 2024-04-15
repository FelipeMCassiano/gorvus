package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func CreateComposeCreateCommand() *cobra.Command {
	var composeTemplate string

	composeCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			workingDir, getWdError := os.Getwd()
			if getWdError != nil {
				fmt.Println(text.FgRed.Sprint("oops! could not get current working directory."))
				os.Exit(1)
			}
			dockerComposePath := path.Join(workingDir, "docker-compose.yml")

			if _, err := os.Stat(dockerComposePath); err == nil {
				fmt.Println(text.FgRed.Sprint("docker-compose.yml already exists. If you want to add a new service use `compose add` command"))
				os.Exit(1)
			}

			prompt := promptui.Select{
				Label: "Select an template",
				Items: []string{"Postgres", "None"},
			}
			_, composeTemplate, _ = prompt.Run()

			if composeTemplate == "None" {
				fmt.Println(text.FgYellow.Sprint("\n No template specified. Creating an empty docker-compose.yml file"))
				os.Create("docker-compose.yml")
				os.Exit(0)
			}

			if err := builders.BuilderComposefile(composeTemplate); err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(text.FgGreen.Sprint("\ndocker-compose.yml created succesfully!"))
		},
	}

	return composeCreateCmd
}
