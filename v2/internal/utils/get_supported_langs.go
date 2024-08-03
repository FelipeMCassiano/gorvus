package utils

func GetSupportedLangs() []string {
	langs := []string{}
	for k := range supportedLangs {
		langs = append(langs, k)
	}
	return langs
}
