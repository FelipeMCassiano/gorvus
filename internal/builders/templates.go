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

func applyTemplate(writer io.Writer, languageTemplate string, data interface{}) {
	tmpl, err := template.New("docker").Parse(languageTemplate)
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
