package dockerfilebuilders

import (
	"embed"
	"errors"

	"github.com/manifoldco/promptui"
)

//go:embed templates/*
var templatesContent embed.FS

type DockerfileData struct {
	EntryFile   string
	Version     string
	ProjectName string
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
