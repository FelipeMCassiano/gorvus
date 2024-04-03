package utils

import "fmt"

func ShowSupportedLangs() {
	langs := supportedLangs()

	fmt.Println("supported languages:")
	for k := range langs {
		fmt.Println("  " + k)
	}
}
