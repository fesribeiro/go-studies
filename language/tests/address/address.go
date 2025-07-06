package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AddressType determines the type of address based on the first word.
// It checks if the first word matches any of the predefined types.
func AddressType(address string) string {
	types := []string{
		"street",
		"avenue",
		"road",
	}

	firstWord := strings.Split(strings.ToLower(address), " ")[0]

	hasValidType := false

	for _, t := range types {
		if strings.EqualFold(firstWord, t) {
			hasValidType = true
			break
		}
	}

	if hasValidType {
		return cases.Title(language.English).String(firstWord)
	}

	return "Unknown"
}
