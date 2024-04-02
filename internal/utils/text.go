package utils

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func NormalizeText(text string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	transformedString, _, err := transform.String(t, text)

	if err != nil {
		return "", err
	}

	return strings.ToLower(transformedString), nil
}
