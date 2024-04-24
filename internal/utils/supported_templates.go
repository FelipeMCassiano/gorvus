package utils

import (
	"github.com/FelipeMCassiano/gorvus/internal/builders/compose"
	"github.com/FelipeMCassiano/gorvus/internal/builders/dockerfile"
)

var supportedLangs = map[string]func(dockerfile.DockerfileData) error{
	"go":            dockerfile.BuildGoDockerfile,
	"rust":          dockerfile.BuildRustDockerfile,
	"node-ts":       dockerfile.BuildTypescriptNodeDockefile,
	"node-js":       dockerfile.BuildJavascriptDockerfile,
	"bun":           dockerfile.BuildBunDockerfile,
	"csharp-dotnet": dockerfile.BuildDotNetDockerfile,
	"java-gradle":   dockerfile.BuilderJavaGradleDockerfile,
}

var supportedComposeTemplates = map[string]func() error{
	"postgres": compose.PostgresBuilderComposefile,
	"mysql":    compose.MysqlBuilderComposefile,
	"mongodb":  compose.MongoDbMBuilderComposefile,
}
