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
