package utils

import "strings"

func FindSimilarLangs(language string) string {
	var similarLangs []string
	supportedLangs := supportedLangs()

	for k := range supportedLangs {
		for _, char := range language {
			if strings.Contains(strings.ToLower(k), string(char)) {
				similarLangs = append(similarLangs, k)
				break
			}
		}
	}
	if len(similarLangs) > 1 {
		return strings.Join(similarLangs, " or ")
	}

	return similarLangs[0]
}
