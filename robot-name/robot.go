package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var names []string

type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getRandomRobotName() string {
	return fmt.Sprintf("%s%03d", randSeq(2), rand.Intn(999))
}

func checkIfExists(name string) bool {
	for _, local := range names {
		if local == name {
			return true
		}
	}
	return false
}

func getUniqueName() string {
	var name string
	for {
		name = getRandomRobotName()
		if checkIfExists(name) {
			continue
		} else {
			return name
		}
	}
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = getUniqueName()
		names = append(names, r.name)
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
