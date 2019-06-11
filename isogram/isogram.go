package isogram

import "strings"

func IsIsogram(input string) bool {
	counts := make(map[string]int)
	for _, char := range input {
		key := strings.ToLower(string(char))
		if key == " " || key == "-" {
			continue
		}
		counts = *setOrIncrement(&counts, key)
	}

	return checkIsogramDict(&counts)
}

func setOrIncrement(counts *map[string]int, key string) *map[string]int {
	dictionary := *counts
	if val, ok := dictionary[key]; ok {
		dictionary[key] = val + 1
	} else {
		dictionary[key] = 1
	}
	return &dictionary
}

func checkIsogramDict(counts *map[string]int) bool {
	flag := true
	for _, value := range *counts {
		flag = flag && (value == 1)
	}
	return flag
}
