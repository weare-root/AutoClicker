package ui

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gotk3/gotk3/glib"
	"github.com/weare-root/AutoClicker/errs"
)

func isHoldingMode() bool {
	rbObj, err := gBuilder.GetObject("rbHold")
	errs.ErrorCheck(err)
	rbHold, err := isRadioButton(rbObj)
	errs.ErrorCheck(err)

	return rbHold.GetActive()
}

// execMainThread is just a wrapper around glib.IdleAdd in case I forget the function name lol
func execMainThread(f interface{}) {
	glib.IdleAdd(f)
}

func getKey() string {
	if rbCustom.GetActive() {
		// custom radio button is selected
		str, err := enCustom.GetText()
		errs.ErrorCheck(err)
		return str
	} else if rbLeft.GetActive() {
		return "left"
	} else if rbMiddle.GetActive() {
		return "center"
	}

	return "right"
}

func getPath(p ...string) string {
	base, err := os.Executable()
	if err != nil {
		base = os.Args[0]
	}
	str, err := filepath.Abs(path.Join(filepath.Dir(base), path.Join(p...)))
	errs.ErrorCheck(err)
	return str
}
