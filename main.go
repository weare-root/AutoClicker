package main

import (
	"log"
	"os"
	// "time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "at.wunderwuzis.AutoClicker"

var (
	gApplication *gtk.Application
	gBuilder     *gtk.Builder
	gWin         *gtk.ApplicationWindow
)

// definitions for the elements
var (
	rbLeft   *gtk.RadioButton
	rbMiddle *gtk.RadioButton
	rbRight  *gtk.RadioButton
	rbCustom *gtk.RadioButton
	rbHold   *gtk.RadioButton

	cbEnabled *gtk.CheckButton

	btKey    *gtk.Button
	btCustom *gtk.Button

	enKey    *gtk.Entry
	enCustom *gtk.Entry

	scCPS *gtk.Scale
)

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	gApplication = application

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		onActivate(application)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		stopClicker()
		log.Println("application shutdown")
	})

	// Launch the application
	application.Run(os.Args)
}

func onActivate(application *gtk.Application) {
	builder, err := gtk.BuilderNewFromFile("./ui/main.glade")
	errorCheck(err)
	gBuilder = builder

	// Get the object with the id of "main_window".
	obj, err := builder.GetObject("main_window")
	errorCheck(err)

	// add signals
	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"activation_btn_clicked": activationBtnClicked,
		"custom_btn_clicked":     customBtnClicked,
		"rbCustom_toggled":       rbCustomToggled,
	}

	builder.ConnectSignals(signals)

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	win, err := isApplicationWindow(obj)
	errorCheck(err)
	gWin = win

	//add styling
	provider, err := gtk.CssProviderNew()
	errorCheck(err)
	err = provider.LoadFromPath("./ui/main.css")
	errorCheck(err)
	screen := win.GetScreen()
	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	win.AddEvents(int(gdk.KEY_PRESS_MASK))

	application.AddWindow(win)
	// Show the Window and all of its components.
	win.Show()

	getWidgets()

	// go startClicker(10, 15, 0.5, 1000, "left")
	// time.Sleep(time.Second * 3)
	// stopClicker()
}

func getWidgets() {
	rbLeft = getRadioButton("rbLeft")
	rbMiddle = getRadioButton("rbMiddle")
	rbRight = getRadioButton("rbRight")
	rbCustom = getRadioButton("rbCustom")
	rbHold = getRadioButton("rbHold")

	cbEnabled = getCheckButton("cbEnabled")

	btKey = getButton("btKey")
	btCustom = getButton("btCustom")

	enKey = getEntry("enKey")
	enCustom = getEntry("enCustom")

	scCPS = getScale("scCPS")
}
