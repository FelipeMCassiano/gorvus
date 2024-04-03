package utils

import (
	"fmt"
	"os"
)

func VerifyIfLangIsSupported(language string) {
	langs := supportedLangs()

	_, ok := langs[language]

	if !ok {
		fmt.Printf("The language %s is not supported", language)
		os.Exit(1)
	}
}
