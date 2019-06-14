// Package clock implements time calculations without date
package clock

import "fmt"

type Clock struct {
	minutes int
}

func New(hour, minute int) Clock {
	var clock Clock
	clock.minutes = (hour*60 + minute) % 1440
	if clock.minutes < 0 {
		clock.minutes += 1440
	}
	return clock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.minutes/60, clock.minutes%60)
}

func (clock Clock) Add(minutes int) Clock {
	return New(0, clock.minutes+minutes)
}

func (clock Clock) Subtract(minutes int) Clock {
	return New(0, clock.minutes-minutes)
}
