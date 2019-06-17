package pangram

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const numLetters = 26

func IsPangram(input string) bool {
	runes := []rune(strings.ToLower(input))
	presence := make(map[rune]int)
	for i := 0; i < utf8.RuneCountInString(input); i++ {
		if unicode.IsLetter(runes[i]) {
			presence[runes[i]]++
		}
	}
	return len(presence) == numLetters
}
