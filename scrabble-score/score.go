package scrabble

import (
	"strings"
)

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(input string) int {
	total := 0
	for _, character := range input {
		for characters, score := range scores {
			if strings.Contains(characters, strings.ToUpper(string(character))) {
				total += score
			}
		}
	}

	return total
}
