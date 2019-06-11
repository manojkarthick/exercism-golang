package reverse

import (
	"unicode/utf8"
)

func Reverse(input string) string {
	s := []rune(input)
	for i, j := 0, utf8.RuneCountInString(input)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
