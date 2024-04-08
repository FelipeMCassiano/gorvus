package utils

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
)

var supportedLangs = map[string]bool{
	"go":      true,
	"rust":    true,
	"node-ts": true,
	"node-js": true,
	"bun-tsx": true,
}

func ShowSupportedLangs() {
	fmt.Println(text.FgGreen.Sprint("Supported languages") + ":")
	for k := range supportedLangs {
		fmt.Println("  " + k)
	}
}
