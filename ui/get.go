package ui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/weare-root/AutoClicker/errs"
)

func getLabel(id string) *gtk.Label {
	lb, err := isLabel(getObject(id))
	errs.ErrorCheck(err)
	return lb
}

func getButton(id string) *gtk.Button {
	bt, err := isButton(getObject(id))
	errs.ErrorCheck(err)
	return bt
}

func getRadioButton(id string) *gtk.RadioButton {
	rb, err := isRadioButton(getObject(id))
	errs.ErrorCheck(err)
	return rb
}

func getEntry(id string) *gtk.Entry {
	en, err := isEntry(getObject(id))
	errs.ErrorCheck(err)
	return en
}

func getCheckButton(id string) *gtk.CheckButton {
	cb, err := isCheckButton(getObject(id))
	errs.ErrorCheck(err)
	return cb
}

func getScale(id string) *gtk.Scale {
	sc, err := isScale(getObject(id))
	errs.ErrorCheck(err)
	return sc
}

func getObject(id string) glib.IObject {
	obj, err := gBuilder.GetObject(id)
	errs.ErrorCheck(err)
	return obj
}

func getAdjustment(id string) *gtk.Adjustment {
	ad, err := isAdjustment(getObject(id))
	errs.ErrorCheck(err)
	return ad
}

func getWindow(id string) *gtk.Window {
	win, err := isWindow(getObject(id))
	errs.ErrorCheck(err)
	return win
}

func getApplicationWindow(id string) *gtk.ApplicationWindow {
	win, err := isApplicationWindow(getObject(id))
	errs.ErrorCheck(err)
	return win
}

func getScrolledWindow(id string) *gtk.ScrolledWindow {
	win, err := isScrolledWindow(getObject(id))
	errs.ErrorCheck(err)
	return win
}

func getComboBoxText(id string) *gtk.ComboBoxText {
	cb, err := isComboBoxText(getObject(id))
	errs.ErrorCheck(err)
	return cb
}
