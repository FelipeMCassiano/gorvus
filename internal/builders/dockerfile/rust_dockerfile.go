package dockerfile

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
)

func RustDockerFileBuilder(input DockerfileData) error {
	if len(input.ProjectName) == 0 {
		pN, err := setProjectName()
		if err != nil {
			os.Exit(1)
		}
		input.ProjectName = pN
	}

	if len(input.EntryFile) >= 1 {
		fmt.Println(text.FgYellow.Sprintf("This language doens't needs to specify the EntryFile"))
		os.Exit(1)
	}
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

	datafile, err := os.ReadFile("templates/dockerfile/rust_dockerfile.tmpl")
	if err != nil {
		fmt.Println("Error parsing Dockerfile template:", err)
		return err
	}

	input.Version = rustVersion

	file, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error creating Dockerfile:", err)
		return err
	}
	defer file.Close()

	builders.ApplyTemplate(file, string(datafile), input)

	return nil
}
