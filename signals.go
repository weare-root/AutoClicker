package main

import (
	"log"
)

var (
	running = false
	before  = ""
)

func activationBtnClicked() {
	running = false
	//activation_btn_clicked
	log.Println("asd")
}

func customBtnClicked() {
	// EvChan := hook.Start()
	// defer hook.End()

	// for ev := range EvChan {
	// 	fmt.Println("hook: ", ev)
	// }
}

// code for listening to the group of radio buttons
func rbLeftToggled() {
	clicksChanged("rbLeft")
}
func rbMiddleToggled() {
	clicksChanged("rbMiddle")
}
func rbRightToggled() {
	clicksChanged("rbRight")
}
func rbCustomToggled() {
	clicksChanged("rbCustom")
}

func clicksChanged(new string) {
	btObj, err := gBuilder.GetObject("btCustom")
	errorCheck(err)
	bt, err := isButton(btObj)
	errorCheck(err)
	if new == "rbCustom" {
		bt.SetSensitive(true)
	} else {
		bt.SetSensitive(false)
	}
}

// code for listening the hold and switch buttons at the top
func rbHoldToggled() {

}
