package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// activationBtnClicked handles the event when the button for the activation key is clicked
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

// customBtnClicked handles the event when the button for the custom keyboard press is clicked
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

// rbCustomToggled listen to the custom radio button
func rbCustomToggled(rbCustom *gtk.RadioButton) {
	btCustom.SetSensitive(rbCustom.GetActive())
}
