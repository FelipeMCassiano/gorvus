package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func CreateComposeAddNetCommand() *cobra.Command {
	var networkNameFlag string
	var networkDriverFlag string
	var nameDockerNetworkFlag string

	composeNetworkCmd := &cobra.Command{
		Use:   "add-net",
		Short: "Adds a new network into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(networkNameFlag) == 0 {
				prompt := promptui.Prompt{
					Label:    "Network Name",
					Validate: validatePrompt,
				}
				name, _ := prompt.Run()
				networkNameFlag = name
			}

			network := Network{
				Driver: networkDriverFlag,
				Name:   nameDockerNetworkFlag,
			}

			if len(network.Driver) == 0 {
				prompt := promptui.Prompt{
					Label:    "Network Driver",
					Validate: validatePrompt,
				}
				driver, _ := prompt.Run()
				network.Driver = driver
			}

			if len(network.Name) == 0 {
				prompt := promptui.Prompt{
					Label:    "Network container name",
					Validate: validatePrompt,
				}
				nameNetwork, _ := prompt.Run()
				network.Name = nameNetwork

			}

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

			newCompose, err := networkAdd(&composeYml, networkNameFlag, network)
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
			fmt.Println(text.FgGreen.Sprint("Network added to docker-compose.yml succesfully!"))
		},
	}

	composeNetworkCmd.Flags().StringVarP(&networkNameFlag, "name", "n", "", "Set the network name")
	composeNetworkCmd.Flags().StringVarP(&networkDriverFlag, "driver", "d", "", "Set the network driver")
	composeNetworkCmd.Flags().StringVarP(&nameDockerNetworkFlag, "name-docker", "x", "", "Set the Docker network name")

	return composeNetworkCmd
}

func networkAdd(compose *DockerCompose, networkName string, network Network) (*DockerCompose, error) {
	newCompose := compose
	if newCompose.Networks == nil {
		newCompose.Networks = make(Networks)
	}

	for inComposeNetworkName := range newCompose.Networks {
		if inComposeNetworkName == networkName {
			return nil, fmt.Errorf("%s is conflicting with a service with same name", networkName)
		}
	}

	newCompose.Networks[networkName] = network

	return newCompose, nil
}
