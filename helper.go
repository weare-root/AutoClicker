package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
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
			activationBtn, _ = toChange.GetText()
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

// execMainThread is just a wrapper around glib.IdleAdd in case I forget the function name lol
func execMainThread(f interface{}) {
	glib.IdleAdd(f)
}

func getKey() string {
	if rbCustom.GetActive() {
		// custom radio button is selected
		str, err := enCustom.GetText()
		errorCheck(err)
		return str
	} else if rbLeft.GetActive() {
		return "left"
	} else if rbMiddle.GetActive() {
		return "middle"
	}

	return "right"
}
