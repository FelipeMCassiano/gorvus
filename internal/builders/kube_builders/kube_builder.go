package kubebuilders

import "embed"

type KubeDatabuilder struct {
	ApiVersion string
	Name       string
	AppName    string
	Replicas   string
	Image      string
	Ports      string
}

//go:embed templates/*
var templatesContent embed.FS

// TODO: Soon!
// type KubeData struct {
// 	ApiVersion string
// 	Kind       string
// 	Metadata   Metadata
// 	Spec       Spec
// }

// type Metadata struct {
// 	Name   string
// 	Labels labels
// }

// type labels struct {
// 	App string
// }

// type Spec struct {
// 	Replicas string
// 	Selector selector
// 	Template template
// 	Spec     containers
// }

// type selector struct {
// 	MatchLabels matchlabels
// }

// type matchlabels struct {
// 	App string
// }

// type template struct {
// 	Metadata Metadata
// }

// type containers struct {
// 	Name  string
// 	Image string
// 	Ports ports
// }
// type ports struct {
// 	ContainerPort string
// }
