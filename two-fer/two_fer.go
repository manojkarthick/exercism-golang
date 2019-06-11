package twofer

import "fmt"

// Two-fer implementation function
func ShareWith(name string) string {
	var x string
	switch {
	case name == "":
		x = "you"
	default:
		x = name
	}
	return fmt.Sprintf("One for %s, one for me.", x)
}
