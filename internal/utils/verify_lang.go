package utils

import (
	"fmt"
	"os"
)

func VerifyIfLangIsSupported(language string) {
	_, ok := supportedLangs[language]

	if !ok {
		fmt.Printf("The language '%s' is not supported\n", language)

		similarLangs := FindSimilarLangs(language)
		if similarLangs != "" {
			fmt.Printf("Did you mean %v?\n", similarLangs)
		}

		os.Exit(1)
	}
}
