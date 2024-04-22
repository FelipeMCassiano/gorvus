package composebuilders

import (
	"embed"
)

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

//go:embed templates/*
var templatesContent embed.FS
