package templates

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"
)

func GoDockerfile(projectName string) error {
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
	dockerFileContent := `
FROM golang:{{.Version}}

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o {{.ProjectName}}

CMD ["./{{.ProjectName}}"]
    `

	tmpl, err := template.New("Dockerfile").Parse(dockerFileContent)
	if err != nil {
		fmt.Println("Error parsing Dockerfile template:", err)
		return err
	}

	file, err := os.Create("Dockerfile")
	if err != nil {
		fmt.Println("Error creating Dockerfile:", err)
		return err
	}

	defer file.Close()

	data := struct {
		Version     string
		ProjectName string
	}{
		ProjectName: projectName,
		Version:     string(goVersion),
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing Dockerfile template:", err)
		return err
	}

	return nil
}
