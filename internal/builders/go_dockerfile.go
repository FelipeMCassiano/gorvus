package builders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"
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

	tmpl, err := template.ParseFiles("internal/templates/go-dockerfile.tmpl")
	if err != nil {
		fmt.Println("Error parsing Dockerfile template:", err)
		return err
	}

	data := dockerfileData{
		ProjectName: projectName,
		Version:     string(goVersion),
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error creating Dockerfile:", err)
		return err
	}
	defer file.Close()

	applyTemplate(file, tmpl, data)

	return nil
}
