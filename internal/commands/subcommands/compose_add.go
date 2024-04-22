package subcommands

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func CreateComposeAddCommand() *cobra.Command {
	var serviceNameFlag string
	var serviceImageFlag string
	var servicePortsFlag []string
	var serviceNetworksFlags []string
	var serviceHostnameFlag string

	composeAddCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds a new service into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			envs := viper.GetStringMapString("envs")

			if len(serviceNameFlag) == 0 {

				prompt := promptui.Prompt{
					Label:    "Service name",
					Validate: validatePrompt,
				}
				sN, err := prompt.Run()
				if err != nil {
					os.Exit(1)
				}

				serviceNameFlag = sN

			}

			service := Service{
				Image:       serviceImageFlag,
				Hostname:    serviceHostnameFlag,
				Environment: envs,
				Networks:    serviceNetworksFlags,
				Ports:       servicePortsFlag,
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

			if composeYml.Version == "" {

				prompt := promptui.Select{
					Label: "It seems like your docker-compose file does not have a version defined. Would you like to define one?",
					Items: []string{"yes", "no"},
				}

				_, answer, _ := prompt.Run()

				if answer == "yes" {
					promptVersion := promptui.Prompt{
						Label:    "Type the disired version",
						Validate: validatePrompt,
					}
					version, _ := promptVersion.Run()
					composeYml.Version = version

				}
			}
			newCompose, addServiceError := composeAdd(&composeYml, serviceNameFlag, service)
			if addServiceError != nil {
				fmt.Println(text.FgRed.Sprint(addServiceError))
				return

			}

			newComposeYmlAsBytes, marshalError := yaml.Marshal(newCompose)
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

	return composeAddCmd
}

func composeAdd(compose *DockerCompose, serviceName string, service Service) (*DockerCompose, error) {
	newCompose := compose

	if newCompose.Services == nil {
		newCompose.Services = make(map[string]*Service)
	}

	newservice := setServiceSettings(&service)

	for inComposeServiceName := range newCompose.Services {
		if inComposeServiceName == serviceName {
			return nil, fmt.Errorf("%s is conflicting with a service with same name", serviceName)
		}
	}
	newCompose.Services[serviceName] = newservice

	return newCompose, nil
}

func setServiceSettings(service *Service) *Service {
	data := service
	if len(data.Image) == 0 {
		imagePrompt := promptui.Prompt{
			Label:    "Image",
			Validate: validatePrompt,
		}
		image, _ := imagePrompt.Run()
		data.Image = image
	}
	if len(data.Hostname) == 0 {

		hostnamePrompt := promptui.Prompt{
			Label:    "Hostname",
			Validate: validatePrompt,
		}
		hostname, _ := hostnamePrompt.Run()
		data.Hostname = hostname
	}
	if len(data.Environment) == 0 {
		for {
			promptKey := promptui.Prompt{
				Label:    "Enter a key for the Environment map (or 'stop' to finish)",
				Validate: validatePrompt,
			}
			key, err := promptKey.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			if key == "stop" {
				break
			}

			promptValue := promptui.Prompt{
				Label:    fmt.Sprintf("Enter a value for the key '%s'", key),
				Validate: validatePrompt,
			}
			value, err := promptValue.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			data.Environment[key] = value
		}
	}
	if len(data.Ports) == 0 {
		for {
			promptPort := promptui.Prompt{
				Label:    "Enter a port for the Ports  (or 'stop' to finish)",
				Validate: validatePrompt,
			}
			port, err := promptPort.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			if port == "stop" {
				break
			}

			data.Ports = append(data.Ports, port)
		}
	}
	if len(data.Networks) == 0 {
		for {
			promptNetwork := promptui.Prompt{
				Label:    "Enter a network for Networks (or 'stop' to finish) ",
				Validate: validatePrompt,
			}

			network, err := promptNetwork.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			if network == "stop" {
				break
			}

			data.Networks = append(data.Networks, network)
		}
	}
	return data
}

func validatePrompt(input string) error {
	if len(input) < 1 {
		return errors.New("This field is required")
	}
	return nil
}
