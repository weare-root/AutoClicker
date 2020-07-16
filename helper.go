package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
	hook "github.com/robotn/gohook"
)

var (
	listening = false

)

func isHoldingMode() bool {
	rbObj, err := gBuilder.GetObject("rbHold")
	errorCheck(err)
	rbHold, err := isRadioButton(rbObj)
	errorCheck(err)

	return rbHold.GetActive()
}

func listenToButtons(button *gtk.Button, toChange *gtk.Entry) {
	listening = true
	button.SetLabel("Drück ESC um aufzuhören")
	channel := hook.Start()

	for ev := range channel {
		if !listening {
			hook.End()
			execMainThread(func() {
				button.SetLabel("Taste auswählen")
			})
			break
		}
		if ev.Kind == hook.KeyUp {
			// https://github.com/robotn/gohook/blob/master/tables.go
			// information where I got the codes from
			if ev.Rawcode == keytoraw["escape"] {
				hook.End()
				listening = false
				execMainThread(func() {
					button.SetLabel("Taste auswählen")
				})
				break
			} else if ev.Rawcode == keytoraw["enter"] ||
				ev.Rawcode == keytoraw["shift"] ||
				ev.Rawcode == keytoraw["ctrl"] {
				continue
			}
			execMainThread(func() {
				toChange.SetText(raw2key[ev.Rawcode])
			})
		}
	}
}

func execMainThread(f interface{}) {
	glib.IdleAdd(f)
}
