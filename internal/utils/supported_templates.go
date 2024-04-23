package utils

import (
	composebuilders "github.com/FelipeMCassiano/gorvus/internal/builders/compose_builders"
	dockerfilebuilders "github.com/FelipeMCassiano/gorvus/internal/builders/dockerfile_builders"
)

var supportedLangs = map[string]func(dockerfilebuilders.DockerfileData) error{
	"go":            dockerfilebuilders.BuildGoDockerfile,
	"rust":          dockerfilebuilders.BuildRustDockerfile,
	"node-ts":       dockerfilebuilders.BuildTypescriptNodeDockefile,
	"node-js":       dockerfilebuilders.BuildJavascriptDockerfile,
	"bun":           dockerfilebuilders.BuildBunDockerfile,
	"csharp-dotnet": dockerfilebuilders.BuildDotNetDockerfile,
	"java-gradle":   dockerfilebuilders.BuilderJavaGradleDockerfile,
}

var supportedComposeTemplates = map[string]func() error{
	"postgres": composebuilders.PostgresBuilderComposefile,
	"myqsl":    composebuilders.MysqlBuilderComposefile,
}
