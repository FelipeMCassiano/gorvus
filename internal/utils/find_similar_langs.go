package utils

import (
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func findSimilarLangs(language string) string {
	var similarLangs []string

	threshold := 2

	for k := range supportedLangs {
		original := k
		if strings.Contains(k, "-") {
			k = strings.ReplaceAll(k, "-", "")
		}

		distance := levenshtein.DistanceForStrings([]rune(language), []rune(k), levenshtein.DefaultOptions)
		if distance <= threshold {
			similarLangs = append(similarLangs, original)
		}
	}
	if len(similarLangs) > 1 {
		return strings.Join(similarLangs, " or ")
	}

	if len(similarLangs) == 0 {
		return ""
	}

	return similarLangs[0]
}
