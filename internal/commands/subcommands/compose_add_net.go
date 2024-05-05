package subcommands

import (
	"fmt"

	"github.com/FelipeMCassiano/gorvus/internal/builders/compose"
	"github.com/FelipeMCassiano/gorvus/internal/utils"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func CreateComposeAddNetCommand() *cobra.Command {
	var networkNameFlag string
	var networkDriverFlag string
	var nameDockerNetworkFlag string

	composeNetworkCmd := &cobra.Command{
		Use:   "add-net",
		Short: "Adds a new network into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(networkNameFlag) == 0 && len(networkDriverFlag) == 0 && len(nameDockerNetworkFlag) == 0 {
				promptName := promptui.Prompt{
					Label:    "Network Name",
					Validate: validatePrompt,
				}
				name, _ := promptName.Run()
				promptDriver := promptui.Prompt{
					Label:    "Network Driver",
					Validate: validatePrompt,
				}
				driver, _ := promptDriver.Run()
				promptNameNetwork := promptui.Prompt{
					Label:    "Network container name",
					Validate: validatePrompt,
				}
				nameNetwork, _ := promptNameNetwork.Run()

				network := compose.Network{
					Driver: driver,
					Name:   nameNetwork,
				}
				composeYml, dockerComposeFileInfo, dockerComposePath, err := utils.GetDockerComposePath()
				if err != nil {
					fmt.Println(text.FgRed.Sprint(err))
					return
				}

				newCompose, err := networkAdd(&composeYml, name, network)
				if err != nil {
					fmt.Println(text.FgRed.Sprint(err))
					return

				}

				if err := utils.WriteDockerCompose(newCompose, dockerComposePath, dockerComposeFileInfo); err != nil {
					fmt.Println(text.FgRed.Sprint(err))
					return
				}

				fmt.Println(text.FgGreen.Sprint("Network added to docker-compose.yml succesfully!"))

				return
			}

			network := compose.Network{
				Driver: networkDriverFlag,
				Name:   nameDockerNetworkFlag,
			}
			if len(networkNameFlag) == 0 {
				fmt.Println(text.FgYellow.Sprint("You must define network name. Use '--name' or '-n"))
				return
			}

			if len(network.Driver) == 0 {
				fmt.Println(text.FgYellow.Sprint("You must define network driver. Use '--driver' or '-d"))
				return
			}

			if len(network.Name) == 0 {
				fmt.Println(text.FgYellow.Sprint("You must define network docker name. Use '--name-docker' or '-x"))

				return
			}

			composeYml, dockerComposeFileInfo, dockerComposePath, err := utils.GetDockerComposePath()
			if err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return
			}

			newCompose, err := networkAdd(&composeYml, networkNameFlag, network)
			if err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return

			}

			if err := utils.WriteDockerCompose(newCompose, dockerComposePath, dockerComposeFileInfo); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return
			}

			fmt.Println(text.FgGreen.Sprint("Network added to docker-compose.yml succesfully!"))
		},
	}

	composeNetworkCmd.Flags().StringVarP(&networkNameFlag, "name", "n", "", "Set the network name")
	composeNetworkCmd.Flags().StringVarP(&networkDriverFlag, "driver", "d", "", "Set the network driver")
	composeNetworkCmd.Flags().StringVarP(&nameDockerNetworkFlag, "name-docker", "x", "", "Set the Docker network name")

	return composeNetworkCmd
}

func networkAdd(composeYml *compose.DockerCompose, networkName string, network compose.Network) (*compose.DockerCompose, error) {
	newCompose := composeYml
	if newCompose.Networks == nil {
		newCompose.Networks = make(compose.Networks)
	}

	if _, ok := newCompose.Networks[networkName]; ok {
		return nil, fmt.Errorf("%s is conflicting with a service with same name", networkName)
	}

	newCompose.Networks[networkName] = network

	return newCompose, nil
}
