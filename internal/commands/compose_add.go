package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

<<<<<<< HEAD:internal/commands/composeAdd.go
type DockerCompose struct {
	Version  string
	Services map[string]Service
}

type Service struct {
	Image       string
	Environment map[string]string
	Ports       []string
	Networks    []string
}

func CreateComposeCommand() *cobra.Command {
=======
func compose() *cobra.Command {
>>>>>>> main:internal/commands/compose_add.go
	var serviceNameFlag string
	var serviceImageFlag string
	var servicePortsFlag string

	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manages current directory's docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("use add or remove to manage your docker-compose with gorvus.")
		},
	}

	composeAddCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds a new service into docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if len(serviceNameFlag) == 0 {
				fmt.Println("\n You must specify the name of the service, use `--name` or `-n`")
				cmd.Help()
				os.Exit(1)
			}
			if len(serviceImageFlag) == 0 {
				fmt.Println("\n You must specify the image of the service, use `--serviceimage` or `-i`")
				cmd.Help()
				os.Exit(1)
			}

			envs := viper.GetStringMapString("envs")

			workingDir, getWdError := os.Getwd()
			if getWdError != nil {
				fmt.Println("oops! could not get current working directory.")
				os.Exit(1)
			}

			dockerComposePath := path.Join(workingDir, "docker-compose.yml")
			dockerComposeFileInfo, statComposeError := os.Stat(dockerComposePath)
			if statComposeError != nil {
				fmt.Println("for some reason, it failed to read docker-compose.yml file.")
				os.Exit(1)
			}

			// todo fallback to empty composeYml
			dockerComposeFileContents, readComposeError := os.ReadFile(dockerComposePath)
			if readComposeError != nil {
				fmt.Println("for some reason, it failed to read docker-compose.yml file.")
				os.Exit(1)
			}

			var composeYml DockerCompose

			yamlParseError := yaml.Unmarshal(dockerComposeFileContents, &composeYml)
			if yamlParseError != nil {
				fmt.Println("can't manage docker-compose.yml, the contents of the file are invalid.")
				os.Exit(1)
			}

			// TODO add flag and creation for networks

			service := Service{
				Image:       serviceImageFlag,
				Environment: envs,

				Networks: []string{
					"Networks",
				},
				Ports: []string{
					servicePortsFlag,
				},
			}

			//! composeYml will be mutated
			if addServiceError := composeAdd(&composeYml, serviceNameFlag, service); addServiceError != nil {
				fmt.Println(addServiceError)
				return
			}

			// reupdate yml file in disk
			newComposeYmlAsBytes, marshalError := yaml.Marshal(composeYml)
			if marshalError != nil {
				fmt.Println("can't manage docker-compose.yml, the contents of the file are invalid.")
				return
			}

			os.WriteFile(dockerComposePath, newComposeYmlAsBytes, dockerComposeFileInfo.Mode())
			fmt.Println("service added to docker-compose.yml")
		},
	}

	composeAddCmd.Flags().StringVarP(&serviceNameFlag, "name", "n", "", "sets the service name in docker-compose")
	composeAddCmd.Flags().StringVarP(&serviceImageFlag, "image", "i", "", "sets the image in docker-compose")
	composeAddCmd.Flags().StringVarP(&servicePortsFlag, "ports", "p", "", "sets the port in service in docker-compose")
	composeAddCmd.Flags().StringToString("envs", map[string]string{}, "sets the environments in docker-compose")
	viper.BindPFlag("envs", composeAddCmd.Flags().Lookup("envs"))

	composeCmd.AddCommand(composeAddCmd)

	return composeCmd
}

<<<<<<< HEAD:internal/commands/composeAdd.go
func ComposeAdd(compose *DockerCompose, serviceName string, service Service) error {
=======
func composeAdd(compose *map[string]interface{}, serviceName string, service map[string]interface{}) error {
>>>>>>> main:internal/commands/compose_add.go
	// todo check for version?
	// is compose["services"] uninitialized? (kinda hacky, but it settles for now)
	// if (*compose)["services"] == nil {
	// 	(*compose)["services"] = make(map[string]interface{})
	// }

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
