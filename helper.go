package main

import (
	"fmt"

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
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		hook.End()
		button.SetLabel("Taste auswählen")
	})
	channel := hook.Start()

	for ev := range channel {
		if !listening {
			hook.End()
			button.SetLabel("Taste auswählen")
			break
		}
		if ev.Kind == hook.KeyUp {
			fmt.Println(ev.Keychar)
			fmt.Println(ev.Rawcode)
			fmt.Println(string(ev.Rawcode))
			char := string(ev.Rawcode)[0]
			if 'a' <= char && char <= 'z' {
				fmt.Println("no way")
			}
		}
	}
}
