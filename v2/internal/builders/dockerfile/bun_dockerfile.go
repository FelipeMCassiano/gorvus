package dockerfile

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/FelipeMCassiano/gorvus/v2/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
)

func BunDockerFileBuilder(input DockerfileData, outpath string) error {
	if len(input.EntryFile) == 0 {
		eF, err := setEntryfile()
		if err != nil {
			os.Exit(1)
		}
		input.EntryFile = eF

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

	file, err := builders.CreateFile(outpath, "Dockerfile")
	if err != nil {
		return fmt.Errorf("failed to creating Dockerfile: %s", err.Error())
	}

	defer file.Close()

	builders.ApplyTemplate(file, string(datafile), input)

	return nil
}
