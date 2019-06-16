package letter

import (
	"sync"
	"unicode/utf8"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Using mutex + wait groups
func WaitGroupConcurrentFrequency(inputs []string) FreqMap {
	f := FreqMap{}
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	for _, input := range inputs {
		wg.Add(utf8.RuneCountInString(input))
		for _, character := range input {
			go func(c rune) {
				defer wg.Done()
				mutex.Lock()
				f[c]++
				mutex.Unlock()
			}(character)
		}
	}
	wg.Wait()
	return f
}

func ConcurrentFrequency(inputs []string) FreqMap {
	f := FreqMap{}
	c := make(chan FreqMap)
	for _, input := range inputs {
		go func(input string) { c <- Frequency(input) }(input)
	}

	for range inputs {
		for k, v := range <-c {
			f[k] += v
		}
	}
	return f
}
