package utils

func GetComposeTemplates(template string) func(string) error {
	builder := supportedComposeTemplates[template]
	return builder
}
