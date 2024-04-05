package utils

import "fmt"

var supportedLangs = map[string]bool{
	"go":   true,
	"rust": true,
}

func ShowSupportedLangs() {
	fmt.Println("supported languages:")
	for k := range supportedLangs {
		fmt.Println("  " + k)
	}
}
