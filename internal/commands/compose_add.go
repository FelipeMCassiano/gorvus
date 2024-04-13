package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Service struct {
	Image       string            `yaml:"image"`
	Hostname    string            `yaml:"hostname"`
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
	var servicePortsFlag []string
	var serviceNetworksFlags []string
	var serviceHostnameFlag string

	var networkNameFlag string
	var networkDriverFlag string
	var nameDockerNetworkFlag string

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

	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manages current directory's docker-compose.yml",
	}

	composeCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(composeTemplateFlag) == 0 {
				fmt.Println(text.FgYellow.Sprint("\n No template specified. Creating an empty docker-compose.yml file"))
				os.Create("docker-compose.yml")
				os.Exit(0)
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

			builders.BuilderComposefile(input, composeTemplateFlag)
			fmt.Println(text.FgGreen.Sprint("docker-compose.yml created succesfully!"))
		},
	}

	composeNetworkCmd := &cobra.Command{
		Use:   "add-net",
		Short: "Adds a new network into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(networkNameFlag) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the network name, use `--name` or `-n`"))
				cmd.Help()
				os.Exit(1)
			}

			if len(networkDriverFlag) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the network driver, use `--driver` or `-d`"))
				cmd.Help()
				os.Exit(1)
			}

			if len(nameDockerNetworkFlag) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the network docker network, use `--name-docker` or `-N`"))
				cmd.Help()
				os.Exit(1)
			}

			network := Network{
				Driver: networkDriverFlag,
				Name:   nameDockerNetworkFlag,
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

			if err := networkAdd(&composeYml, networkNameFlag, network); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				return

			}

			newComposeYmlAsBytes, marshalError := yaml.Marshal(composeYml)
			if marshalError != nil {
				fmt.Println(text.FgRed.Sprint("can't manage docker-compose.yml, the contents of the file are invalid."))
				return
			}

			os.WriteFile(dockerComposePath, newComposeYmlAsBytes, dockerComposeFileInfo.Mode())
			fmt.Println(text.FgGreen.Sprint("Network added to docker-compose.yml succesfully!"))
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
			if len(serviceHostnameFlag) == 0 {
				fmt.Println(text.FgRed.Sprint("\n You must specify the hostname of the service, use `--hostname` or `-h`"))
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
				Hostname:    serviceHostnameFlag,
				Environment: envs,
				Networks:    serviceNetworksFlags,
				Ports:       servicePortsFlag,
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
			fmt.Println(text.FgGreen.Sprint("service added to docker-compose.yml succesfully!"))
		},
	}

	composeAddCmd.Flags().StringVarP(&serviceNameFlag, "service", "s", "", "sets the service name in docker-compose")
	composeAddCmd.Flags().StringVarP(&serviceImageFlag, "image", "i", "", "sets the service image in docker-compose")
	composeAddCmd.Flags().StringSliceVarP(&servicePortsFlag, "ports", "p", []string{}, "sets the service port in service in docker-compose")
	composeAddCmd.Flags().StringToStringP("envs", "e", map[string]string{}, "sets an service environment variable in docker-compose")
	viper.BindPFlag("envs", composeAddCmd.Flags().Lookup("envs"))
	composeAddCmd.Flags().StringSliceVarP(&serviceNetworksFlags, "networks", "n", []string{}, "sets the service network in docker-compose")
	composeAddCmd.Flags().StringVarP(&serviceHostnameFlag, "hostname", "o", "", "sets the service hostname in docker-compose")

	composeAddCmd.MarkFlagRequired("service")
	composeAddCmd.MarkFlagRequired("image")

	composeNetworkCmd.Flags().StringVarP(&networkNameFlag, "name", "n", "", "Set the network name")
	composeNetworkCmd.Flags().StringVarP(&networkDriverFlag, "driver", "d", "", "Set the network driver")
	composeNetworkCmd.Flags().StringVarP(&nameDockerNetworkFlag, "name-docker", "x", "", "Set the Docker network name")
	composeNetworkCmd.MarkFlagRequired("name")
	composeNetworkCmd.MarkFlagRequired("driver")
	composeNetworkCmd.MarkFlagRequired("name-docker")

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

	composeCmd.AddCommand(composeAddCmd)
	composeCmd.AddCommand(composeNetworkCmd)
	composeCmd.AddCommand(composeCreateCmd)

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
