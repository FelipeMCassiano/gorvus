package builders

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/manifoldco/promptui"
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

func validateProjectName(input string) error {
	if len(input) < 1 {
		return errors.New("The project name is required")
	}

	return nil
}

func validateEntryfile(input string) error {
	if len(input) < 1 {
		return errors.New("The entry file is required")
	}

	return nil
}

func setProjectName() (string, error) {
	var projectName string
	prompt := promptui.Prompt{
		Label:    "Type your project name",
		Validate: validateProjectName,
	}
	projectName, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return projectName, nil
}

func setEntryfile() (string, error) {
	var entryfile string
	prompt := promptui.Prompt{
		Label:    "Type your entry file",
		Validate: validateEntryfile,
	}
	entryfile, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return entryfile, nil
}

func validatePrompt(input string) error {
	if len(input) < 1 {
		return errors.New("This field is required")
	}
	return nil
}
