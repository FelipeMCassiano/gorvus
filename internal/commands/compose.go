package commands

import (
	"github.com/spf13/cobra"
)

type Service struct {
	Image       string            `yaml:"image"`
	Hostname    string            `yaml:"hostname"`
	Environment map[string]string `yaml:"environment"`
	Ports       []string          `yaml:"ports"`
	Networks    []string          `yaml:"networks"`
}

type Network struct {
	Driver string `yaml:"driver"`
	Name   string `yaml:"name"`
}

type Networks map[string]Network

type DockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Networks Networks           `yaml:"networks"`
}

func CreateComposeCommand() *cobra.Command {
	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manages current directory's docker-compose.yml",
	}
	composeCmd.AddCommand(CreateComposeAddCommand())
	composeCmd.AddCommand(CreateComposeAddNetCommand())
	composeCmd.AddCommand(CreateComposeCreateCommand())

	return composeCmd
}
