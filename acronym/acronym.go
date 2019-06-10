// Package to implement the "acronym" challenge of Exercism go track.
package acronym

import (
	"strings"
	"unicode"
)

/*
Challenge:
- Convert a phrase to its acronym.
- Techies love their TLA (Three Letter Acronyms)!
- Help generate some jargon by writing a program that converts a long name like Portable Network Graphics to its acronym (PNG).
 */

// Abbreviates the given string
func Abbreviate(input string) string {
	var abbrevs []string

	parts := strings.Split(input, " ")
	for _, part := range parts {
		subparts := strings.Split(part, "-")
		for _, subpart := range subparts {
			for _, char := range subpart {
				if unicode.IsLetter(char) {
					abbrevs = append(abbrevs, string(unicode.ToUpper(char)))
					break
				}
			}

		}
	}

	return strings.Join(abbrevs, "")
}
