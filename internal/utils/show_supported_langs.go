package utils

import (
	"fmt"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/jedib0t/go-pretty/v6/text"
)

var supportedLangs = map[string]func(builders.DockerfileData) error{
	"go":      builders.BuildGoDockerfile,
	"rust":    builders.BuildRustDockerfile,
	"node-ts": builders.BuildTypescriptNodeDockefile,
	"node-js": builders.BuildJavascriptDockerfile,
	"bun":     builders.BuildBunDockerfile,
}

func ShowSupportedLangs() {
	fmt.Println(text.FgGreen.Sprint("Supported languages") + ":")
	for k := range supportedLangs {
		fmt.Println("  " + k)
	}
}
