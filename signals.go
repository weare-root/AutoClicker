package main

import (
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func activationBtnClicked() {
	stopClicker()
	if !listening {
		shouldListen = true
		listeningBtn = btKey
		listeningEn = enKey
		execMainThread(func() {
			btKey.SetLabel("ESC zum abbrechen")
		})
	} else {
		shouldListen = false
		execMainThread(func() {
			btKey.SetLabel("Taste auswählen")
		})
	}
}

func customBtnClicked() {
	stopClicker()
	if !listening {
		shouldListen = true
		listeningBtn = btCustom
		listeningEn = enCustom
		execMainThread(func() {
			btCustom.SetLabel("ESC zum abbrechen")
		})
	} else {
		shouldListen = false
		execMainThread(func() {
			btCustom.SetLabel("Taste auswählen")
		})
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
