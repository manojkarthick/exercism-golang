package luhn

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func reverse(input string) string {
	s := []rune(input)
	for i, j := 0, utf8.RuneCountInString(input)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func HasAllNumbers(remark string) bool {
	flag := true
	for _, character := range remark {
		flag = flag && unicode.IsNumber(character)
	}
	return flag
}

func Valid(input string) bool {
	input = reverse(strings.ReplaceAll(strings.TrimSpace(input), " ", ""))
	length := utf8.RuneCountInString(input)
	if length <= 1 || !HasAllNumbers(input) {
		return false
	}
	sum := luhnSum(input) % 10
	if sum != 0 {
		return false
	}
	return true

}

func luhnSum(input string) int64 {

	var total int64
	for index, character := range input {
		digit, err := strconv.ParseInt(string(character), 10, 0)
		if err != nil {
			log.Fatalf("Unable to convert string to integer: %v", err)
		}
		if index%2 == 1 {
			if 2*digit > 9 {
				total += (2 * digit) - 9
			} else {
				total += (2 * digit)
			}
		} else {
			total += digit
		}
	}
	return total
}
