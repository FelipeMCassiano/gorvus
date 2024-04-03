package templates

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"
)

func RustDockerfile(projectName string) error {
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

	dockerFileContent := `
FROM rust:{{.Version}}-buster as builder

WORKDIR /app

COPY . .

RUN cargo build --release

FROM debian:buster-slim

WORKDIR /usr/local/bin

COPY --from=builder /app/target/release/{{.ProjectName}} .

RUN apt-get update && apt install -y openssl

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

	data := dataDockerfile{
		ProjectName: projectName,
		Version:     rustVersion,
	}

	if err := tmpl.Execute(file, data); err != nil {
		fmt.Println("Error executing Dockerfile template:", err)
		return err
	}
	return nil
}
