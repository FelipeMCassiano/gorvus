package utils

import "fmt"

var supportedLangs = map[string]bool{
	"go":      true,
	"rust":    true,
	"node-ts": true,
	"node-js": true,
	"bun-tsx": true,
}

func ShowSupportedLangs() {
	fmt.Println("supported languages:")
	for k := range supportedLangs {
		fmt.Println("  " + k)
	}
}
