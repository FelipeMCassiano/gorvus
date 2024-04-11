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

type DockerfileData struct {
	EntryFile   string
	Version     string
	ProjectName string
}

type ComposeData struct {
	Version      string
	ImageVersion string
	DbName       string
	DbUser       string
	DbPass       string
	Ports        string
	Cpu          string
	Memory       string
	NetworkName  string
}

func applyTemplate(writer io.Writer, Template string, data interface{}) {
	tmpl, err := template.New("dockerfile").Parse(Template)
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
