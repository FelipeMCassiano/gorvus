package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/FelipeMCassiano/gorvus/internal/utils"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func generateDockerfile() *cobra.Command {
	var projectName string
	var language string
	var listLanguages bool
	var entryFile string

	cmd := &cobra.Command{
		Use:     "create-dockerfile",
		Short:   "Create Dockerfile based on input language and project name",
		Aliases: []string{"gend", "generate-dockerfile"},
		Run: func(cmd *cobra.Command, args []string) {
			if listLanguages {
				utils.ShowSupportedLangs()
				os.Exit(0)
			}

			if len(language) == 0 {
				fmt.Println(text.FgYellow.Sprint("> You must specify the language, use --language or -l"))
				os.Exit(1)
			}

			input := builders.DockerfileData{
				ProjectName: projectName,
				EntryFile:   entryFile,
			}

			builder := utils.VerifyIfLangIsSupported(strings.ToLower(language))

			if err := builder(input); err != nil {
				fmt.Println(text.FgRed.Sprintf("error: %s", err.Error()))
				os.Exit(1)

			}
			fmt.Println(text.FgGreen.Sprint("Dockerfile created succesfully!"))
		},
	}

	cmd.Flags().StringVarP(&projectName, "project-name", "p", "", "Define project name")
	cmd.Flags().StringVarP(&language, "language", "l", "", "Define template language")
	cmd.Flags().BoolVarP(&listLanguages, "list-languages", "s", false, "Gives a list with the supported languages")
	cmd.Flags().StringVarP(&entryFile, "entry-file", "e", "", "Define entry file")

	return cmd
}
