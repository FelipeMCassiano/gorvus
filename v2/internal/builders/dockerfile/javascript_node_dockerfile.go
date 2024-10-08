package dockerfile

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/FelipeMCassiano/gorvus/v2/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
)

func JavascriptDockerFileBuilder(input DockerfileData, outpath string) error {
	if len(input.EntryFile) == 0 {
		eF, err := setEntryfile()
		if err != nil {
			os.Exit(1)
		}
		input.EntryFile = eF
	}

	if strings.Contains(input.EntryFile, ".ts") {
		fmt.Println(text.FgYellow.Sprint("> Only allowed .js files"))

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

	datafile, err := templatesContent.ReadFile("templates/javascript_node_dockerfile.tmpl")
	if err != nil {
		return err
	}

	file, err := builders.CreateFile(outpath, "Dockerfile")
	if err != nil {
		return fmt.Errorf("failed to creating Dockerfile: %s", err.Error())
	}

	input.Version = nodeVersion

	defer file.Close()

	builders.ApplyTemplate(file, string(datafile), input)

	return nil
}
