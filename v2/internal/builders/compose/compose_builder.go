package compose

import "embed"

//go:embed templates/*
var templatesContent embed.FS

type ComposeData struct {
	ImageVersion string
	DbName       string
	DbUser       string
	DbRootPass   string
	DbPass       string
	Restart      string
	Ports        string
	Cpu          string
	Memory       string
	NetworkName  string
}

type Service struct {
	Image       string            `yaml:"image"`
	Hostname    string            `yaml:"hostname"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Ports       []string          `yaml:"ports,omitempty"`
	Networks    []string          `yaml:"networks,omitempty"`
}

type Network struct {
	Driver string `yaml:"driver"`
	Name   string `yaml:"name"`
}

type Networks map[string]Network

type DockerCompose struct {
	Services map[string]Service `yaml:"services"`
	Networks Networks           `yaml:"networksk,omitempty"`
}
