package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/FelipeMCassiano/gorvus/internal/builders/compose"
	"gopkg.in/yaml.v3"
)

func GetDockerComposePath(outpath string) (compose.DockerCompose, fs.FileInfo, string, error) {
	var workingDir string
	if outpath != "" {
		if _, err := os.Stat(outpath); err != nil && os.IsNotExist(err) {
			return compose.DockerCompose{}, nil, "", err
		}

		workingDir = outpath
	} else {
		wD, getWdError := os.Getwd()
		if getWdError != nil {
			err := fmt.Errorf("oops! could not get current working directory.")
			return compose.DockerCompose{}, nil, "", err
		}
		workingDir = wD
	}

	dockerComposePath := path.Join(workingDir, "docker-compose.yml")
	dockerComposeFileInfo, statComposeError := os.Stat(dockerComposePath)
	if statComposeError != nil {
		err := fmt.Errorf("for some reason, it failed to read docker-compose.yml file.")
		return compose.DockerCompose{}, nil, "", err
	}

	dockerComposeFileContents, readComposeError := os.ReadFile(dockerComposePath)
	if readComposeError != nil {
		err := fmt.Errorf("for some reason, it failed to read docker-compose.yml file.")
		return compose.DockerCompose{}, nil, "", err
	}
	var composeYml compose.DockerCompose

	yamlParseError := yaml.Unmarshal(dockerComposeFileContents, &composeYml)
	if yamlParseError != nil {
		err := fmt.Errorf("can't manage docker-compose.yml, the contents of the file are invalid.")
		return compose.DockerCompose{}, nil, "", err
	}

	return composeYml, dockerComposeFileInfo, dockerComposePath, nil
}
