package builders 

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var rustDockerfile = `FROM rust:{{.Version}}-buster as builder

WORKDIR /app

COPY . .

RUN cargo build --release

FROM debian:buster-slim

WORKDIR /usr/local/bin

COPY --from=builder /app/target/release/{{.ProjectName}} .

RUN apt-get update && apt install -y openssl

CMD ["./{{.ProjectName}}"]
`

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

	applyTemplate(file, rustDockerfile, data)

	return nil
}
