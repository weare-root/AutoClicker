package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

var toggled bool = false

func doClick(ratio float32, timespan int, mb string) {
	robotgo.MouseToggle("down", mb)
	time.Sleep(time.Duration(int(float32(timespan)*ratio)) * time.Millisecond)
	robotgo.MouseToggle("up", mb)
	time.Sleep(time.Duration(int(float32(timespan)*(1-ratio))) * time.Millisecond)
}

/*
* Starts the clicking process
* mb = (left|center|right)
* lower = min click amount
* upper = max click amount
* timespan = how long for all clicks to complete (ms)
* ratio = how long holding down mb and how long waiting for next click
 */
func startClicker(lower float32, upper float32, ratio float32, timespan int, mb string) {
	toggled = true
OUTER:
	for toggled {
		cps := int(lower + rand.Float32()*(upper-lower))
		for i := 0; i < cps; i++ {
			if !toggled {
				break OUTER
			}
			doClick(ratio, timespan/cps, mb)
		}
	}
}

func stopClicker() {
	toggled = false
}

func toggle() bool {
	toggled = !toggled
	return toggled
}
