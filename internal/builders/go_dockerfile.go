package builders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	_ "embed"
)

func BuildGoDockerfile(projectName string) error {
	cmd := exec.Command("go", "version")
	goVersionOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	versionPattern := `go version go(\d+\.\d+\.\d+)`
	versionRegex := regexp.MustCompile(versionPattern)

	matches := versionRegex.FindStringSubmatch(string(goVersionOutput))
	if len(matches) < 2 {
		return fmt.Errorf("failed to extract Go version number")
	}
	goVersion := matches[1]
	datafile, err := templatesContent.ReadFile("templates/go_dockerfile.tmpl")
	if err != nil {
		return err
	}

	data := dockerfileData{
		ProjectName: projectName,
		Version:     string(goVersion),
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("Error creating Dockerfile: %s", err.Error())
	}
	defer file.Close()

	applyTemplate(file, string(datafile), data)

	return nil
}
