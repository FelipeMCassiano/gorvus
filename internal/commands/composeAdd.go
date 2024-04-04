package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func CreateComposeCommand() *cobra.Command {
	var serviceNameFlag string

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
			if len(args) == 0 {
				fmt.Println("invalid arguments given, the add command accepts a single argument that is the image to be added.")
				fmt.Println("tip: you can also set a custom service name by using the --name flag")
				return
			}

			workingDir, getWdError := os.Getwd()
			if getWdError != nil {
				fmt.Println("oops! could not get current working directory.")
				return
			}

			dockerComposePath := path.Join(workingDir, "docker-compose.yml")
			dockerComposeFileInfo, statComposeError := os.Stat(dockerComposePath)
			if statComposeError != nil {
				fmt.Println("for some reason, it failed to read docker-compose.yml file.")
				return
			}

			// todo fallback to empty composeYml
			dockerComposeFileContents, readComposeError := os.ReadFile(dockerComposePath)
			if readComposeError != nil {
				fmt.Println("for some reason, it failed to read docker-compose.yml file.")
				return
			}

			composeYml := make(map[string]interface{})
			yamlParseError := yaml.Unmarshal(dockerComposeFileContents, composeYml)
			if yamlParseError != nil {
				fmt.Println("can't manage docker-compose.yml, the contents of the file are invalid.")
				return
			}

			selectedImage := args[0]

			// use image name as defaults if service name is not set
			if serviceNameFlag == "" {
				serviceNameFlag = selectedImage
			}

			// todo use templating file
			service := map[string]interface{}{
				"name":  serviceNameFlag,
				"image": selectedImage,
			}

			//! composeYml will be mutated
			if addServiceError := ComposeAdd(&composeYml, serviceNameFlag, service); addServiceError != nil {
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

	composeAddCmd.PersistentFlags().StringVarP(&serviceNameFlag, "name", "n", "", "sets the service name in docker-compose")

	composeCmd.AddCommand(composeAddCmd)

	return composeCmd
}

func ComposeAdd(compose *map[string]interface{}, serviceName string, service map[string]interface{}) error {
	// todo check for version?
	composeServices := (*compose)["services"].(map[string]interface{})

	// search for conflicting service names
	for inComposeServiceName := range composeServices {
		if inComposeServiceName == serviceName {
			return fmt.Errorf("%s is conflicting with a service with same name", serviceName)
		}
	}

	// todo maybe prevent this side effect by returning new yml?
	// add requested service into compose services
	composeServices[serviceName] = service

	return nil
}
