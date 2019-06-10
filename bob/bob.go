// Package to implement the "bob" challenge of Exercism go track.
package bob

import (
	"strings"
	"unicode"
)

/*
Challenge
- Bob is a lackadaisical teenager. In conversation, his responses are very limited.
- Bob answers 'Sure.' if you ask him a question.
- He answers 'Whoa, chill out!' if you yell at him.
- He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
- He says 'Fine. Be that way!' if you address him without actually saying anything.
- He answers 'Whatever.' to anything else.
Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.
 */

func HasLetters(remark string) bool{
	flag := false
	for _, character := range remark {
		flag = flag || unicode.IsLetter(character)
	}
	return flag
}


// Makes Bob respond based on the remarks from Alice
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	question := strings.HasSuffix(remark,"?")
	yell := strings.Compare(remark, strings.ToUpper(remark)) == 0 && HasLetters(remark)
	empty := strings.Compare(strings.TrimSpace(remark), "") == 0

	switch {
	case question && yell:
		return "Calm down, I know what I'm doing!"
	case yell:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	case empty:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
