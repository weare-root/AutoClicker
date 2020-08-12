package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/weare-root/AutoClicker/autoclicker"
	"github.com/weare-root/AutoClicker/errs"
	"github.com/weare-root/AutoClicker/ui"
)

const appID = "at.wunderwuzis.AutoClicker"

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errs.ErrorCheck(err)

	// set dark mode :)
	settings, err := gtk.SettingsGetDefault()
	errs.LogOnError(err, "could not get settings")
	if err == nil {
		settings.Set("gtk-application-prefer-dark-theme", true)
	}

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activation")
		ui.OnActivate(application)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		autoclicker.StopClicker()
		log.Println("application shutdown")
	})

	// Launch the application
	application.Run(os.Args)
}
