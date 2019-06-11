package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(input int) string {
	var buffer bytes.Buffer
	if input%3 == 0 {
		buffer.WriteString("Pling")
	}
	if input%5 == 0 {
		buffer.WriteString("Plang")
	}
	if input%7 == 0 {
		buffer.WriteString("Plong")
	}
	if input%3 != 0 && input%5 != 0 && input%7 != 0 {
		buffer.WriteString(strconv.Itoa(input))
	}
	return buffer.String()
}
