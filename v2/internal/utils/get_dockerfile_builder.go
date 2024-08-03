package utils

import (
	"fmt"
	"os"

	"github.com/FelipeMCassiano/gorvus/internal/builders/dockerfile"
	"github.com/jedib0t/go-pretty/v6/text"
)

func GetDockerfileBuilder(language string) func(dockerfile.DockerfileData, string) error {
	builder, ok := supportedLangs[language]

	if !ok {
		fmt.Println(
			text.FgRed.Sprintf("The language '%s' is not supported.", language),
		)

		similarLangs := findSimilarLangs(language)
		if similarLangs != "" {
			fmt.Printf("Did you mean %v?\n", similarLangs)
		}

		os.Exit(1)
	}

	return builder
}
