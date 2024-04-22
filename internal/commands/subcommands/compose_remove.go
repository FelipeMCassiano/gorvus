package subcommands

import (
	"fmt"
	"os"
	"path"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func CreateComposeRemoveCommand() *cobra.Command {
	var networkFlag string
	var serviceFlag string
	composeDeleteCmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove services or networks in docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			workingDir, getWdError := os.Getwd()
			if getWdError != nil {
				fmt.Println(text.FgRed.Sprint("oops! could not get current working directory."))
				os.Exit(1)
			}

			dockerComposePath := path.Join(workingDir, "docker-compose.yml")
			dockerComposeFileInfo, statComposeError := os.Stat(dockerComposePath)
			if statComposeError != nil {
				fmt.Println(text.FgRed.Sprint("for some reason, it failed to read docker-compose.yml file."))
				os.Exit(1)
			}

			dockerComposeFileContents, readComposeError := os.ReadFile(dockerComposePath)
			if readComposeError != nil {
				fmt.Println(text.FgRed.Sprint("for some reason, it failed to read docker-compose.yml file."))
				os.Exit(1)
			}

			var composeYml DockerCompose

			yamlParseError := yaml.Unmarshal(dockerComposeFileContents, &composeYml)
			if yamlParseError != nil {
				fmt.Println(text.FgRed.Sprint("can't manage docker-compose.yml, the contents of the file are invalid."))
				os.Exit(1)
			}

			newCompose, err := remove(&composeYml, serviceFlag, networkFlag)
			if err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return
			}

			newComposeYmlAsBytes, marshalError := yaml.Marshal(newCompose)
			if marshalError != nil {
				fmt.Println(text.FgRed.Sprint("can't manage docker-compose.yml, the contents of the file are invalid."))
				return
			}

			os.WriteFile(dockerComposePath, newComposeYmlAsBytes, dockerComposeFileInfo.Mode())
			fmt.Println(text.FgGreen.Sprint("Removed succesfully!"))
		},
	}

	composeDeleteCmd.Flags().StringVarP(&networkFlag, "network", "n", "", "Specify the network name to remove")
	composeDeleteCmd.Flags().StringVarP(&serviceFlag, "service", "s", "", "Speficy the service name to remove")

	return composeDeleteCmd
}

func remove(compose *DockerCompose, serviceName string, networkName string) (*DockerCompose, error) {
	newCompose := compose

	if serviceName != "" {
		if len(newCompose.Services) == 0 {
			return nil, fmt.Errorf("Cannot remove '%s'. Seems like you have no service defined yet. If would like to create one consider use 'gorvus compose add' command", serviceName)
		}
		for inComposeSerivceName := range newCompose.Services {
			if inComposeSerivceName != serviceName {
				return nil, fmt.Errorf("Cannot remove '%s'. This service doens't exists", serviceName)
			}
		}
		delete(compose.Services, serviceName)
	}

	if networkName != "" {
		if len(newCompose.Networks) == 0 {
			return nil, fmt.Errorf("Cannot remove %s. Seems like you have no network defined yet. If would like to create one consider use 'gorvus compose add' command", networkName)
		}
		for inComposeNetworkName := range newCompose.Networks {
			if inComposeNetworkName != networkName {
				return nil, fmt.Errorf("Cannot remove '%s'. This network doens't exists", networkName)
			}
		}
		delete(compose.Networks, networkName)
	}
	return newCompose, nil
}
