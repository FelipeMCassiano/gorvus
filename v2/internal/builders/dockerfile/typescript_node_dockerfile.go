package dockerfile

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
)

func TypescriptDockerFileBuilder(input DockerfileData, outpath string) error {
	if len(input.EntryFile) == 0 {
		eF, err := setEntryfile()
		if err != nil {
			os.Exit(1)
		}
		input.EntryFile = eF

	}

	if strings.Contains(input.EntryFile, ".js") {
		fmt.Println(text.FgYellow.Sprint("> Only allowed .ts files"))

		os.Exit(1)
	}

	if len(input.ProjectName) >= 1 {
		fmt.Println(text.FgYellow.Sprintf("This language doens't needs to specify the Project Name"))
		os.Exit(1)

	}
	cmd := exec.Command("node", "-v")
	nodeVersionOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	versionPattern := `v(\d+\.\d+\.\d+)`
	versionRegex := regexp.MustCompile(versionPattern)

	matches := versionRegex.FindStringSubmatch(string(nodeVersionOutput))
	if len(matches) < 2 {
		return fmt.Errorf("failed to extract Node version number")
	}

	nodeVersion := matches[1]

	datafile, err := templatesContent.ReadFile("templates/typescript_node_dockerfile.tmpl")
	if err != nil {
		return err
	}

	input.EntryFile = strings.TrimSuffix(input.EntryFile, filepath.Ext(input.EntryFile))
	input.Version = nodeVersion

	file, err := builders.CreateFile(outpath, "Dockerfile")
	if err != nil {
		return fmt.Errorf("failed to creating Dockerfile: %s", err.Error())
	}

	defer file.Close()

	builders.ApplyTemplate(file, string(datafile), input)

	return nil
}
