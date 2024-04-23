package utils

func GetSupportedComposeTemplates() []string {
	templateSupported := []string{}
	for k := range supportedComposeTemplates {
		templateSupported = append(templateSupported, k)
	}
	templateSupported = append(templateSupported, "None")

	return templateSupported
}
