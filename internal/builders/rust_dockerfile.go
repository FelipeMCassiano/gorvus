package builders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func BuildRustDockerfile(projectName string) error {
	cmd := exec.Command("cargo", "-V")
	rustVersionOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	versionPattern := `cargo (\d+\.\d+\.\d+) `

	versionRegex := regexp.MustCompile(versionPattern)

	matches := versionRegex.FindStringSubmatch(string(rustVersionOutput))
	if len(matches) < 2 {
		return fmt.Errorf("failed to extract Rust version number")
	}
	rustVersion := matches[1]

	datafile, err := templatesContent.ReadFile("templates/rust_dockerfile.tmpl")
	if err != nil {
		fmt.Println("Error parsing Dockerfile template:", err)
		return err
	}

	data := dockerfileData{
		ProjectName: projectName,
		Version:     rustVersion,
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error creating Dockerfile:", err)
		return err
	}
	defer file.Close()

	applyTemplate(file, string(datafile), data)

	return nil
}
