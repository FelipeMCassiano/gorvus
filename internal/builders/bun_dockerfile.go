package builders

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

func BuildBunDockerfile(input DockerfileData) error {
	if len(input.EntryFile) == 0 {
		fmt.Println(text.FgYellow.Sprint("> You must specify the entry file, use `--entry-file` or `-e`"))
		os.Exit(1)
	}
	if len(input.ProjectName) >= 1 {
		fmt.Println(text.FgYellow.Sprintf("This language doens't needs to specify the Project Name"))
		os.Exit(1)

	}

	if !strings.Contains(input.EntryFile, ".ts") && !strings.Contains(input.EntryFile, ".js") {
		fmt.Println(text.FgYellow.Sprint("> You must choose between files types .js or .ts"))
		os.Exit(1)
	}

	cmd := exec.Command("bun", "-v")
	bunVersionOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	datafile, err := templatesContent.ReadFile("templates/tsx_bun_dockerfile.tmpl")
	if err != nil {
		return err
	}

	bunVersion := strings.TrimSpace(string(bunVersionOutput))

	input.Version = bunVersion

	file, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("failed to creating Dockerfile: %s", err.Error())
	}

	defer file.Close()

	applyTemplate(file, string(datafile), input)

	return nil
}
