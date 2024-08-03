package utils

import (
	"github.com/FelipeMCassiano/gorvus/v2/internal/builders/compose"
	"github.com/FelipeMCassiano/gorvus/v2/internal/builders/dockerfile"
)

var supportedLangs = map[string]func(dockerfile.DockerfileData, string) error{
	"go":            dockerfile.GoDockerFileBuilder,
	"rust":          dockerfile.RustDockerFileBuilder,
	"node-ts":       dockerfile.TypescriptDockerFileBuilder,
	"node-js":       dockerfile.JavascriptDockerFileBuilder,
	"bun":           dockerfile.BunDockerFileBuilder,
	"csharp-dotnet": dockerfile.DotNetDockerFileBuilder,
	"java-gradle":   dockerfile.JavaGradleDockerFileBuilder,
}

var supportedComposeTemplates = map[string]func(string) error{
	"postgres": compose.PostgreSQLComposeFileBuilder,
	"mysql":    compose.MySQLComposeFileBuilder,
	"mongodb":  compose.MongoDBComposeFileBuilder,
}
