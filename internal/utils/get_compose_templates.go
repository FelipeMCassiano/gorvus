package utils

func GetComposeTemplates(template string) func() error {
	builder := supportedComposeTemplates[template]
	return builder
}
