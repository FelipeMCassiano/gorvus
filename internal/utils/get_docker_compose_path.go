package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/FelipeMCassiano/gorvus/internal/builders/compose"
	"github.com/jedib0t/go-pretty/v6/text"
	"gopkg.in/yaml.v3"
)

func GetDockerComposePath() (compose.DockerCompose, fs.FileInfo, string, error) {
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
	var composeYml compose.DockerCompose

	yamlParseError := yaml.Unmarshal(dockerComposeFileContents, &composeYml)
	if yamlParseError != nil {
		fmt.Println(text.FgRed.Sprint("can't manage docker-compose.yml, the contents of the file are invalid."))
	}

	return composeYml, dockerComposeFileInfo, dockerComposePath, nil
}
