package builders

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

type dockerfileData struct {
	Version     string
	ProjectName string
}

func applyTemplate(writer io.Writer, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(writer, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}
}
