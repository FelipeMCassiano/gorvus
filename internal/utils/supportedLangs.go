package utils

type supportedLanguages map[string]bool

func supportedLangs() supportedLanguages {
	languagesSuported := map[string]bool{
		"go":   true,
		"rust": true,
	}

	return languagesSuported
}
