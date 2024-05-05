package utils

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/FelipeMCassiano/gorvus/internal/builders/compose"
	"gopkg.in/yaml.v3"
)

func WriteDockerCompose(newCompose *compose.DockerCompose, dockerComposePath string, dockerComposeFileInfo fs.FileInfo) error {
	newComposeYmlAsBytes, marshalError := yaml.Marshal(newCompose)
	if marshalError != nil {
		return fmt.Errorf("can't manage docker-compose.yml, the contents of the file are invalid.")
	}

	if err := os.WriteFile(dockerComposePath, newComposeYmlAsBytes, dockerComposeFileInfo.Mode()); err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
