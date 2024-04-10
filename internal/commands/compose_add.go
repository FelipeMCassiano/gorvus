package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Service struct {
	Image       string            `yaml:"image"`
	Environment map[string]string `yaml:"environment"`
	Ports       []string          `yaml:"ports"`
	Networks    []string          `yaml:"networks"`
}

type Network struct {
	Driver string `yaml:"driver"`
	Name   string `yaml:"name"`
}

type Networks map[string]Network

type DockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Networks Networks           `yaml:"networks"`
}

func CreateComposeCommand() *cobra.Command {
	var serviceNameFlag string
	var serviceImageFlag string
	var servicePortsFlag string
	var serviceNetworksFlags []string

	var networkName string
	var networkDriver string
	var nameDockerNetwork string

	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manages current directory's docker-compose.yml",
	}

	composeNetworkCmd := &cobra.Command{
		Use:   "add-net",
		Short: "Adds a new network into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(networkName) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the network name, use `--name` or `-n`"))
				cmd.Help()
				os.Exit(1)
			}

			if len(networkDriver) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the network driver, use `--driver` or `-d`"))
				cmd.Help()
				os.Exit(1)
			}

			if len(nameDockerNetwork) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the network docker network, use `--name-docker` or `-N`"))
				cmd.Help()
				os.Exit(1)
			}

			network := Network{
				Driver: networkDriver,
				Name:   nameDockerNetwork,
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

			// todo fallback to empty composeYml
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

			if err := networkAdd(&composeYml, networkName, network); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return

			}

			newComposeYmlAsBytes, marshalError := yaml.Marshal(composeYml)
			if marshalError != nil {
				fmt.Println(text.FgRed.Sprint("can't manage docker-compose.yml, the contents of the file are invalid."))
				return
			}

			os.WriteFile(dockerComposePath, newComposeYmlAsBytes, dockerComposeFileInfo.Mode())
			fmt.Println(text.FgGreen.Sprint("network added to docker-compose.yml"))
		},
	}

	composeAddCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds a new service into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(serviceNameFlag) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the name of the service, use `--name` or `-n`"))
				cmd.Help()
				os.Exit(1)
			}
			if len(serviceImageFlag) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the image of the service, use `--serviceimage` or `-i`"))
				cmd.Help()
				os.Exit(1)
			}

			envs := viper.GetStringMapString("envs")

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

			// todo fallback to empty composeYml
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

			if composeYml.Version == "" {
				var answer string
				fmt.Println("You want to update version? (y/n) ")
				fmt.Scanln(&answer)
				if answer == "y" {
					var version string
					fmt.Println("Type the desired version: ")
					fmt.Scanln(&version)
					composeYml.Version = version

				}
			}

			// TODO add flag and creation for networks

			service := Service{
				Image:       serviceImageFlag,
				Environment: envs,
				Networks:    serviceNetworksFlags,
				Ports: []string{
					servicePortsFlag,
				},
			}

			//! composeYml will be mutated
			if addServiceError := composeAdd(&composeYml, serviceNameFlag, service); addServiceError != nil {
				fmt.Println(text.FgRed.Sprint(addServiceError))
				return
			}

			// reupdate yml file in disk
			newComposeYmlAsBytes, marshalError := yaml.Marshal(composeYml)
			if marshalError != nil {
				fmt.Println(text.FgRed.Sprint("can't manage docker-compose.yml, the contents of the file are invalid."))
				return
			}

			os.WriteFile(dockerComposePath, newComposeYmlAsBytes, dockerComposeFileInfo.Mode())
			fmt.Println(text.FgGreen.Sprint("service added to docker-compose.yml"))
		},
	}

	composeAddCmd.Flags().StringVarP(&serviceNameFlag, "service", "s", "", "sets the service name in docker-compose")
	composeAddCmd.Flags().StringVarP(&serviceImageFlag, "image", "i", "", "sets the service image in docker-compose")
	composeAddCmd.Flags().StringVarP(&servicePortsFlag, "ports", "p", "", "sets the service port in service in docker-compose")
	composeAddCmd.Flags().StringToStringP("envs", "e", map[string]string{}, "sets an service environment variable in docker-compose")
	viper.BindPFlag("envs", composeAddCmd.Flags().Lookup("envs"))
	composeAddCmd.Flags().StringSliceVarP(&serviceNetworksFlags, "networks", "n", []string{}, "sets the service network in docker-compose")

	composeAddCmd.MarkFlagRequired("service")
	composeAddCmd.MarkFlagRequired("image")

	composeNetworkCmd.Flags().StringVarP(&networkName, "name", "n", "", "Set the network name")
	composeNetworkCmd.Flags().StringVarP(&networkDriver, "driver", "d", "", "Set the network driver")
	composeNetworkCmd.Flags().StringVarP(&nameDockerNetwork, "name-docker", "x", "", "Set the Docker network name")
	composeNetworkCmd.MarkFlagRequired("name")
	composeNetworkCmd.MarkFlagRequired("driver")
	composeNetworkCmd.MarkFlagRequired("name-docker")

	composeCmd.AddCommand(composeAddCmd)
	composeCmd.AddCommand(composeNetworkCmd)

	return composeCmd
}

func networkAdd(compose *DockerCompose, networkName string, network Network) error {
	if compose.Networks == nil {
		compose.Networks = make(Networks)
	}

	for inComposeNetworkName := range compose.Networks {
		if inComposeNetworkName == networkName {
			return fmt.Errorf("%s is conflicting with a service with same name", networkName)
		}
	}

	compose.Networks[networkName] = network

	return nil
}

func composeAdd(compose *DockerCompose, serviceName string, service Service) error {
	// todo check for version?
	// is compose["services"] uninitialized? (kinda hacky, but it settles for now)

	if compose.Services == nil {
		compose.Services = make(map[string]Service)
	}

	// composeServices := (*compose)["services"].(map[string]interface{})

	// search for conflicting service names
	for inComposeServiceName := range compose.Services {
		if inComposeServiceName == serviceName {
			return fmt.Errorf("%s is conflicting with a service with same name", serviceName)
		}
	}
	// todo maybe prevent this side effect by returning new yml?
	// add requested service into compose services
	compose.Services[serviceName] = service

	return nil
}
