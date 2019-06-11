package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of the strings is not equal, cannot calculate hamming distance.")
	}

	var h int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			h++
		}
	}

	return h, nil
}
