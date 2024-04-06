package builders

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func BuildTsxBunDockerfile(entryfile string) error {
	cmd := exec.Command("bun", "-v")
	bunVersionOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	datafile, err := templatesContent.ReadFile("templates/tsx_bun_dockerfile.tmpl")
	if err != nil {
		return err
	}

	bunVersion := strings.TrimSpace(string(bunVersionOutput))

	data := dockerfileData{
		EntryFile: entryfile,
		Version:   bunVersion,
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("failed to creating Dockerfile: %s", err.Error())
	}

	defer file.Close()

	applyTemplate(file, string(datafile), data)

	return nil
}
