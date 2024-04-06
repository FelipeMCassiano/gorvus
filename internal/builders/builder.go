package builders

import (
	"embed"
	"fmt"
	"io"
	"os"
	"text/template"
)

//go:embed templates/*
var templatesContent embed.FS

type dockerfileData struct {
	EntryFile   string
	Version     string
	ProjectName string
}

func applyTemplate(writer io.Writer, languageTemplate string, data interface{}) {
	tmpl, err := template.New("dockerfile").Parse(languageTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}
}
