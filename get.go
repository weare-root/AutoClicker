package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func getButton(id string) *gtk.Button {
	bt, err := isButton(getObject(id))
	errorCheck(err)
	return bt
}

func getRadioButton(id string) *gtk.RadioButton {
	rb, err := isRadioButton(getObject(id))
	errorCheck(err)
	return rb
}

func getEntry(id string) *gtk.Entry {
	en, err := isEntry(getObject(id))
	errorCheck(err)
	return en
}

func getCheckButton(id string) *gtk.CheckButton {
	cb, err := isCheckButton(getObject(id))
	errorCheck(err)
	return cb
}

func getScale(id string) *gtk.Scale {
	sc, err := isScale(getObject(id))
	errorCheck(err)
	return sc
}

func getObject(id string) glib.IObject {
	obj, err := gBuilder.GetObject(id)
	errorCheck(err)
	return obj
}

func getAdjustment(id string) *gtk.Adjustment {
	ad, err := isAdjustment(getObject(id))
	errorCheck(err)
	return ad
}
