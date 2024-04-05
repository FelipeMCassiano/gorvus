package builders

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var goDockerfile = `FROM golang:{{.Version}}

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o {{.ProjectName}}

CMD ["./{{.ProjectName}}"]
`

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

	applyTemplate(file, goDockerfile, data)

	return nil
}
