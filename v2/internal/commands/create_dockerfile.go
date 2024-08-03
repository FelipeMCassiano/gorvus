package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/FelipeMCassiano/gorvus/v2/internal/builders/dockerfile"
	"github.com/FelipeMCassiano/gorvus/v2/internal/utils"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func generateDockerfile() *cobra.Command {
	var projectName string
	var language string
	var listLanguages bool
	var entryFile string
	var outputPath string

	cmd := &cobra.Command{
		Use:     "create-dockerfile",
		Short:   "Create Dockerfile based on input language and project name",
		Aliases: []string{"gend", "generate-dockerfile"},
		Run: func(cmd *cobra.Command, args []string) {
			if outputPath != "" {
				if _, err := os.Stat(outputPath); err != nil && os.IsNotExist(err) {
					fmt.Println(text.FgRed.Sprint(err.Error()))
					os.Exit(1)
				}
			}
			if listLanguages {
				utils.ShowSupportedLangs()
				os.Exit(0)
			}

			if len(language) == 0 {
				prompt := promptui.Select{
					Label: "Select language",
					Items: utils.GetSupportedLangs(),
				}
				_, language, _ = prompt.Run()
			}

			input := dockerfile.DockerfileData{
				ProjectName: projectName,
				EntryFile:   entryFile,
			}

			builder := utils.GetDockerfileBuilder(strings.ToLower(language))

			if err := builder(input, outputPath); err != nil {
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
	cmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Define output path")

	return cmd
}
