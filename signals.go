package main

import (
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func activationBtnClicked() {
	stopClicker()
	if listening {
		btKey.SetLabel("Taste auswählen")
		listening = false
	} else {
		go listenToButtons(btKey, enKey)
	}
}

func customBtnClicked() {
	stopClicker()
	if listening {
		btCustom.SetLabel("Taste auswählen")
		listening = false
	} else {
		go listenToButtons(btCustom, enCustom)
	}
}

// listen to the custom radio button
func rbCustomToggled() {
	btObj, err := gBuilder.GetObject("btCustom")
	errorCheck(err)
	bt, err := isButton(btObj)
	errorCheck(err)

	rbObj, err := gBuilder.GetObject("rbCustom")
	errorCheck(err)
	rb, err := isRadioButton(rbObj)
	errorCheck(err)

	bt.SetSensitive(rb.GetActive())
}

func enableToggled() {
	toggled = false
	enabled = cbEnabled.GetActive()
}

func onDurationInsert(en *gtk.Entry, text string) bool {
	before, _ := en.GetText()
	parsed, err := strconv.ParseInt(before+text, 10, 64)
	if err != nil {
		defer execMainThread(func() {
			en.SetText(before)
		})
	}
	duration = int(parsed)
	return false
}
