package builders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func BuildJavascriptDockerfile(entryfile string) error {
	cmd := exec.Command("node", "-v")
	nodeVersionOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	versionPattern := `v(\d+\.\d+\.\d+)`
	versionRegex := regexp.MustCompile(versionPattern)

	matches := versionRegex.FindStringSubmatch(string(nodeVersionOutput))
	if len(matches) < 2 {
		return fmt.Errorf("failed to extract Node version number")
	}

	nodeVersion := matches[1]

	datafile, err := templatesContent.ReadFile("templates/javascript_node_dockerfile.tmpl")
	if err != nil {
		return err
	}
	data := dockerfileData{
		EntryFile: entryfile,
		Version:   nodeVersion,
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("failed to creating Dockerfile: %s", err.Error())
	}

	defer file.Close()

	applyTemplate(file, string(datafile), data)

	return nil
}
