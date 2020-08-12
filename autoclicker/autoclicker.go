package autoclicker

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var toggled bool = false

// doClick simulates the mouse click
func doClick(ratio float32, timespan int, mb string) {
	robotgo.MouseToggle("down", mb)
	time.Sleep(time.Duration(int(float32(timespan)*ratio)) * time.Millisecond)
	robotgo.MouseToggle("up", mb)
	time.Sleep(time.Duration(int(float32(timespan)*(1-ratio))) * time.Millisecond)
}

// doPress simulates a keyboard press
func doPress(ratio float32, timespan int, key string) {
	robotgo.KeyTap(key)
	time.Sleep(time.Duration(int(float32(timespan)*ratio)) * time.Millisecond)
}

// startMouseClicker starts the clicking process
// mb = (left|center|right)
// lower = min click amount
// upper = max click amount
// timespan = how long for all clicks to complete (ms)
// ratio = how long holding down mb and how long waiting for next click
func startMouseClicker(lower, upper, ratio float32, timespan int, mb string) {
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
		time.Sleep(time.Duration(int(1000-timespan)) * time.Millisecond)
	}
}

// StartClicker starts either the mouse clicker or the keyboard clicker
func StartClicker(lower, upper, ratio float32, timespan int, key string) {
	key = strings.TrimSpace(key)
	if key == "" {
		return
	}
	log.Println("starting clicker")
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
		time.Sleep(time.Duration(int(1000-timespan)) * time.Millisecond)
	}
}

// StopClicker stops the clicker
func StopClicker() {
	if toggled {
		log.Println("Stopping clicker")
	}
	toggled = false
}
