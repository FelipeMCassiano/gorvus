package builders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	_ "embed"

	"github.com/jedib0t/go-pretty/v6/text"
)

func BuildGoDockerfile(input DockerfileData) error {
	if len(input.ProjectName) == 0 {
		pN, err := setProjectName()
		if err != nil {
			os.Exit(1)
		}
		input.ProjectName = pN
	}

	if len(input.EntryFile) >= 1 {
		fmt.Println(text.FgYellow.Sprintf("> This language doens't needs to specify the Entry File"))
		os.Exit(1)
	}

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

	input.Version = goVersion

	file, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("Error creating Dockerfile: %s", err.Error())
	}
	defer file.Close()

	applyTemplate(file, string(datafile), input)

	return nil
}
