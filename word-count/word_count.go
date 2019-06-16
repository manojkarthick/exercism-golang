package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	words := strings.FieldsFunc(strings.ToLower(phrase), split)
	freq := make(Frequency)

	for _, word := range words {
		word = strings.TrimFunc(word, unquote)
		re := regexp.MustCompile("([a-z0-9']+)")
		match := re.FindStringSubmatch(word)
		if len(match) > 0 {
			freq[match[0]]++
		}
	}

	return freq
}

func split(r rune) bool {
	return r == ' ' || r == ','
}

func unquote(r rune) bool {
	return r == '\'' || r == '"'
}
