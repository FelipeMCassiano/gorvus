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
	var composeTemplateFlag string
	var composeVersionFlag string
	var composeImageVersionFlag string
	var composeDbNameFlag string
	var composeUserFlag string
	var composePassFlag string
	var composePortsFlag string
	var composeCpuFlag string
	var composeMemoryFlag string
	var composeNetworkName string
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

			if len(composeTemplateFlag) == 0 {
				prompt := promptui.Select{
					Label: "Select an template",
					Items: []string{"Postgres", "None"},
				}
				_, composeTemplateFlag, _ = prompt.Run()

				if composeTemplateFlag == "None" {
					fmt.Println(text.FgYellow.Sprint("\n No template specified. Creating an empty docker-compose.yml file"))
					os.Create("docker-compose.yml")
					os.Exit(0)
				}

			}

			input := builders.ComposeData{
				Version:      composeVersionFlag,
				ImageVersion: composeImageVersionFlag,
				DbName:       composeDbNameFlag,
				DbUser:       composeUserFlag,
				DbPass:       composePassFlag,
				Ports:        composePortsFlag,
				Cpu:          composeCpuFlag,
				Memory:       composeMemoryFlag,
				NetworkName:  composeNetworkName,
			}

			if err := builders.BuilderComposefile(input, composeTemplateFlag); err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(text.FgGreen.Sprint("docker-compose.yml created succesfully!"))
		},
	}

	composeCreateCmd.Flags().StringVarP(&composeTemplateFlag, "template", "t", "", "defines template")
	composeCreateCmd.Flags().StringVarP(&composeVersionFlag, "version", "v", "", "defines compose version")
	composeCreateCmd.Flags().StringVarP(&composeImageVersionFlag, "image-version", "i", "", "defines image version")
	composeCreateCmd.Flags().StringVarP(&composeDbNameFlag, "db-name", "d", "", "defines db name environment")
	composeCreateCmd.Flags().StringVarP(&composeUserFlag, "user", "u", "", "defines user environment")
	composeCreateCmd.Flags().StringVarP(&composePassFlag, "password", "a", "", "defines password environment")
	composeCreateCmd.Flags().StringVarP(&composePortsFlag, "ports", "p", "", "defines ports")
	composeCreateCmd.Flags().StringVarP(&composeCpuFlag, "cpu", "c", "", "defines cpu deploy resources")
	composeCreateCmd.Flags().StringVarP(&composeMemoryFlag, "memory", "m", "", "defines memory deploy resources")
	composeCreateCmd.Flags().StringVarP(&composeNetworkName, "network-name", "n", "", "defines network name")

	return composeCreateCmd
}
