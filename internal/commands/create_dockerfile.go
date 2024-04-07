package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/FelipeMCassiano/gorvus/internal/utils"
	"github.com/spf13/cobra"
)

func generateDockerfile() *cobra.Command {
	var projectName string
	var language string
	var listLanguages bool
	var entryFile string
	const defaultProjectName = "myproject"

	cmd := &cobra.Command{
		Use:     "create-dockerfile",
		Short:   "Create Dockerfile based on input language and project name",
		Aliases: []string{"gend"},
		Run: func(cmd *cobra.Command, args []string) {
			if listLanguages {
				utils.ShowSupportedLangs()
				os.Exit(0)
			}

			if len(language) == 0 {
				fmt.Println("\n> You must specify the language, use --language or -l")
				os.Exit(1)

			}

			if len(projectName) == 0 {
				fmt.Println("\n> WARN:")
				fmt.Printf("\n>You haven’t specified the project. If the language requires project name, it will be created with %s", defaultProjectName)
			}

			utils.VerifyIfLangIsSupported(language)

			if strings.ToLower(language) == "go" {
				if err := builders.BuildGoDockerfile(projectName); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}
			}

			if strings.ToLower(language) == "rust" {
				if err := builders.BuildRustDockerfile(projectName); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}
			}

			if strings.ToLower(language) == "node-ts" {
				if len(entryFile) == 0 {
					fmt.Println("\n> You must specify the entry file, use `--entry-file` or `-e`")
					os.Exit(1)
				}
				if err := builders.BuildTypescriptNodeDockefile(entryFile); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}

			}

			if strings.ToLower(language) == "node-js" {
				if len(entryFile) == 0 {
					fmt.Println("\n> You must specify the entry file, use `--entry-file` or `-e`")
					os.Exit(1)
				}
				if err := builders.BuildJavascriptDockerfile(entryFile); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}

			}

			if strings.ToLower(language) == "bun-tsx" {
				if len(entryFile) == 0 {
					fmt.Println("\n> You must specify the entry file, use `--entry-file` or `-e`")
					os.Exit(1)
				}

				if !strings.Contains(entryFile, ".ts") && !strings.Contains(entryFile, ".js") {
					fmt.Println("\n> You must choose between files types .js or .ts")
					os.Exit(1)
				}

				if err := builders.BuildTsxBunDockerfile(entryFile); err != nil {
					fmt.Printf("error: %s", err.Error())
					os.Exit(1)
				}

			}
			fmt.Println("Dockerfile created succesfully")
		},
	}

	cmd.Flags().StringVarP(&projectName, "project-name", "p", defaultProjectName, "Define project name")
	cmd.Flags().StringVarP(&language, "language", "l", "", "Define template language")
	cmd.Flags().BoolVarP(&listLanguages, "list-languages", "s", false, "Gives a list with the supported languages")
	cmd.Flags().StringVarP(&entryFile, "entry-file", "e", "", "Define entry file")

	return cmd
}
