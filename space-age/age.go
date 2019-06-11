package space

import "math"

var orbitalPeriods = map[string]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const secondsInEarthYear = 31557600

func Age(seconds float64, planet string) float64 {
	earthYears := seconds / secondsInEarthYear
	age := earthYears / orbitalPeriods[planet]

	if planet == "Earth" {
		return earthYears
	}
	return math.Round(age*100) / 100
}
