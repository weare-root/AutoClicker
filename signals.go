package main

func activationBtnClicked() {
	stopClicker()
	if listening {
		listening = false
	} else {
		go listenToButtons(btKey, enKey)
	}
}

func customBtnClicked() {
	stopClicker()
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
