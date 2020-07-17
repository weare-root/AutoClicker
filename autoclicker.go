package main

import (
	"log"
	"math/rand"
	"strings"
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

func doPress(ratio float32, timespan int, key string) {
	robotgo.KeyToggle(key, "down")
	time.Sleep(time.Duration(int(float32(timespan)*ratio)) * time.Millisecond)
	robotgo.KeyToggle(key, "up")
	time.Sleep(time.Duration(int(float32(timespan)*(1-ratio))) * time.Millisecond)
}

// startClicker starts the clicking process
// mb = (left|center|right)
// lower = min click amount
// upper = max click amount
// timespan = how long for all clicks to complete (ms)
// ratio = how long holding down mb and how long waiting for next click
func startMouseClicker(lower float32, upper float32, ratio float32, timespan int, mb string) {
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

func startClicker(lower float32, upper float32, ratio float32, timespan int) {
	key := strings.TrimSpace(getKey())
	log.Println(key)
	if key == "" {
		return
	}
	if key == "left" || key == "right" || key == "center" {
		startMouseClicker(lower, upper, ratio, timespan, key)
		return
	}

	toggled = true
OUTER:
	for toggled {
		cps := int(lower + rand.Float32()*(upper-lower))
		for i := 0; i < cps; i++ {
			if !toggled {
				break OUTER
			}
			doPress(ratio, timespan/cps, key)
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
