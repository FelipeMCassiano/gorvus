package dockerfilebuilders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
)

func BuildDotNetDockerfile(input DockerfileData) error {
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

	cmd := exec.Command("dotnet", "--version")
	dotnetVersionOutput, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	versionPattern := `(\d+)\.`
	re := regexp.MustCompile(versionPattern)

	matches := re.FindStringSubmatch(string(dotnetVersionOutput))

	if len(matches) < 2 {
		return fmt.Errorf("failed to extract Dotnet version number")
	}
	input.Version = matches[1]

	datafile, err := templatesContent.ReadFile("templates/dotnet_dockerfile.tmpl")
	if err != nil {
		return err
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		return err
	}

	defer file.Close()

	builders.ApplyTemplate(file, string(datafile), input)

	return nil
}