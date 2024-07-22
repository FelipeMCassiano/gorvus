package builders

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

func ApplyTemplate(writer io.Writer, Template string, data interface{}) {
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

func CreateFile(path string, fileName string) (*os.File, error) {
	return os.Create(filepath.Join(path, fileName))
}
