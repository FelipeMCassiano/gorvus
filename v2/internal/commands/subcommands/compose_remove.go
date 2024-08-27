package subcommands

import (
	"fmt"
	"os"

	"github.com/FelipeMCassiano/gorvus/v2/internal/builders/compose"
	"github.com/FelipeMCassiano/gorvus/v2/internal/utils"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func CreateComposeRemoveCommand() *cobra.Command {
	var networkFlag string
	var serviceFlag string
	composeDeleteCmd := &cobra.Command{
		Use:     "remove",
		Short:   "Remove services or networks in docker-compose.yml",
		Aliases: []string{"rm"},
		Run: func(cmd *cobra.Command, args []string) {
			changedDirectory, err := cmd.Flags().GetString("cd")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			composeYml, dockerComposeFileInfo, dockerComposePath, err := utils.GetDockerComposePath(changedDirectory)
			if err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return
			}
			newCompose, err := remove(&composeYml, serviceFlag, networkFlag)
			if err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return
			}

			if err := utils.WriteDockerCompose(newCompose, dockerComposePath, dockerComposeFileInfo); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return
			}

			fmt.Println(text.FgGreen.Sprint("Removed succesfully!"))
		},
	}

	composeDeleteCmd.Flags().StringVarP(&networkFlag, "network", "n", "", "Specify the network name to remove")
	composeDeleteCmd.Flags().StringVarP(&serviceFlag, "service", "s", "", "Speficy the service name to remove")

	return composeDeleteCmd
}

func remove(compose *compose.DockerCompose, serviceName string, networkName string) (*compose.DockerCompose, error) {
	newCompose := compose

	if serviceName != "" {
		if len(newCompose.Services) == 0 {
			return nil, fmt.Errorf("Cannot remove '%s'. Seems like you have no service defined yet. If would like to create one consider use 'gorvus compose add' command", serviceName)
		}
		if _, ok := newCompose.Services[serviceName]; !ok {
			return nil, fmt.Errorf("Cannot remove '%s'. This service doens't exists", serviceName)
		}
		delete(compose.Services, serviceName)
	}

	if networkName != "" {
		if len(newCompose.Networks) == 0 {
			return nil, fmt.Errorf("Cannot remove %s. Seems like you have no network defined yet. If would like to create one consider use 'gorvus compose add' command", networkName)
		}
		if _, ok := newCompose.Networks[networkName]; !ok {
			return nil, fmt.Errorf("Cannot remove '%s'. This network doens't exists", networkName)
		}
		delete(compose.Networks, networkName)
	}
	return newCompose, nil
}
