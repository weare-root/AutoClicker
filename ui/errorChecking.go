package ui

import (
	"errors"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func isApplicationWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}

func isEntry(obj glib.IObject) (*gtk.Entry, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Entry); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Entry")
}

func isListBox(obj glib.IObject) (*gtk.ListBox, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.ListBox); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.ListBox")
}

func isBox(obj glib.IObject) (*gtk.Box, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Box); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Box")
}

func isLabel(obj glib.IObject) (*gtk.Label, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Label); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Label")
}

func isImage(obj glib.IObject) (*gtk.Image, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Image); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Image")
}

func isScrolledWindow(obj glib.IObject) (*gtk.ScrolledWindow, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.ScrolledWindow); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.ScrolledWindow")
}

func isButton(obj glib.IObject) (*gtk.Button, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Button); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Button")
}
func isRadioButton(obj glib.IObject) (*gtk.RadioButton, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.RadioButton); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.RadioButton")
}

func isScale(obj glib.IObject) (*gtk.Scale, error) {
	// make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Scale); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.scale")
}

func isCheckButton(obj glib.IObject) (*gtk.CheckButton, error) {
	// make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.CheckButton); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.CheckButton")
}

func isAdjustment(obj glib.IObject) (*gtk.Adjustment, error) {
	// make type assertion (as per gtk.go).
	if adjustment, ok := obj.(*gtk.Adjustment); ok {
		return adjustment, nil
	}
	return nil, errors.New("not a *gtk.CheckButton")
}

func isComboBoxText(obj glib.IObject) (*gtk.ComboBoxText, error) {
	// make type assertion (as per gtk.go).
	if combobox, ok := obj.(*gtk.ComboBoxText); ok {
		return combobox, nil
	}
	return nil, errors.New("not a *gtk.ComboBoxText")
}
