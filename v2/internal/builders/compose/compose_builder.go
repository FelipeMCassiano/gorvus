package compose

import "embed"

//go:embed templates/*
var templatesContent embed.FS

type ComposeData struct {
	Version      string
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
