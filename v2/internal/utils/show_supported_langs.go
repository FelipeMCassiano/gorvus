package utils

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
)

func ShowSupportedLangs() {
	fmt.Println(text.FgGreen.Sprint("Supported languages") + ":")
	for k := range supportedLangs {
		fmt.Println("  " + k)
	}
}
