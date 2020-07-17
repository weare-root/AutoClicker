package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	hook "github.com/robotn/gohook"
)

const appID = "at.wunderwuzis.AutoClicker"

var (
	gApplication *gtk.Application
	gBuilder     *gtk.Builder
	gWin         *gtk.ApplicationWindow
)

// variables for runtime
var (
	activationBtn = ""
	duration      = 0
	listening     = false
	listeningBtn  *gtk.Button
	listeningEn   *gtk.Entry
	shouldListen  = false
)

// definitions for the elements
var (
	rbLeft   *gtk.RadioButton
	rbMiddle *gtk.RadioButton
	rbRight  *gtk.RadioButton
	rbCustom *gtk.RadioButton
	rbHold   *gtk.RadioButton

	btKey    *gtk.Button
	btCustom *gtk.Button

	enKey    *gtk.Entry
	enCustom *gtk.Entry

	scCPSLower  *gtk.Scale
	scCPSHigher *gtk.Scale
	scTimespan  *gtk.Scale
	scRatio     *gtk.Scale

	adLower    *gtk.Adjustment
	adHigher   *gtk.Adjustment
	adTimespan *gtk.Adjustment
	adRatio    *gtk.Adjustment
)

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	gApplication = application

	// set dark mode :)
	settings, err := gtk.SettingsGetDefault()
	logOnError(err, "could not get settings")
	if err == nil {
		settings.Set("gtk-application-prefer-dark-theme", true)
	}

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
	log.Println("application activation")
	log.Println("UI is being built")
	builder, err := gtk.BuilderNewFromFile(getPath("ui", "main.glade"))
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
		"on_duration_insert":     onDurationInsert,
	}

	builder.ConnectSignals(signals)
	log.Println("connected signals")

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	win, err := isApplicationWindow(obj)
	errorCheck(err)
	gWin = win

	//add styling
	screen := win.GetScreen()

	// add a gtk theme for Windows
	if runtime.GOOS == "windows" {
		theme, err := gtk.CssProviderNew()
		errorCheck(err)

		err = theme.LoadFromPath(getPath("ui", "win10-theme", "gtk.css"))
		errorCheck(err)

		gtk.AddProviderForScreen(screen, theme, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	}

	// add own customizations (not much lol)
	provider, err := gtk.CssProviderNew()
	errorCheck(err)
	err = provider.LoadFromPath(getPath("ui", "main.css"))
	errorCheck(err)
	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	log.Println("Applied custom styles")

	win.AddEvents(int(gdk.KEY_PRESS_MASK))

	application.AddWindow(win)
	// Show the Window and all of its components.
	win.Show()

	getWidgets()

	go listenToKeyboard()

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

	btKey = getButton("btKey")
	btCustom = getButton("btCustom")

	enKey = getEntry("enKey")
	enCustom = getEntry("enCustom")

	scCPSLower = getScale("scCPSLower")
	scCPSHigher = getScale("scCPSHigher")
	scTimespan = getScale("scTimespan")
	scRatio = getScale("scRatio")

	adLower = getAdjustment("adLower")
	adHigher = getAdjustment("adHigher")
	adTimespan = getAdjustment("adTimespan")
	adRatio = getAdjustment("adRatio")
}

// doTheDirty is the listening function
func listenToKeyboard() {
	isDown := false
	channel := hook.Start()
	lastChanged := getCurrentMillis()
	log.Println("started listening to keyboard")

	for ev := range channel {
		if shouldListen && !listening {
			listening = true
		} else if !shouldListen {
			listening = false
		}

		if listening {
			if ev.Kind == hook.KeyUp {
				if ev.Rawcode == keytoraw["escape"] {
					shouldListen = false
					execMainThread(func() {
						listeningBtn.SetLabel("Taste auswählen")
					})
					continue
				} else if ev.Rawcode == keytoraw["enter"] ||
					ev.Rawcode == keytoraw["shift"] ||
					ev.Rawcode == keytoraw["ctrl"] {
					continue
				}
				activationBtn = raw2key[ev.Rawcode]
				execMainThread(func() {
					listeningEn.SetText(raw2key[ev.Rawcode])
				})
			}
			continue
		}

		if ev.Rawcode == keytoraw[activationBtn] && activationBtn != "" {
			// check if its a hold for the holding mode
			if ev.Kind == hook.KeyHold && isHoldingMode() {
				if !isDown {
					// start the holding if it isnt already started
					log.Println("HOLDING start")
					isDown = true
					go startClicker(float32(adLower.GetValue()), float32(adHigher.GetValue()), float32(adRatio.GetValue()/100.0), int(adTimespan.GetValue()))
				}
			} else if ev.Kind == hook.KeyUp {
				if getCurrentMillis()-lastChanged < 500 {
					continue
				}
				lastChanged = getCurrentMillis()

				if isHoldingMode() {
					log.Println("HOLDING stop")
					// stop when holding mode is on
					isDown = false
					stopClicker()
				} else if !isDown {
					log.Println("SWITCHING start")
					// start when switching mode is on
					isDown = true
					go startClicker(float32(adLower.GetValue()), float32(adHigher.GetValue()), float32(adRatio.GetValue()/100.0), int(adTimespan.GetValue()))
				} else {
					// stop when switching mode is on
					log.Println("SWITCHING stop")
					isDown = false
					stopClicker()
				}
			}

		}
	}
}

func getCurrentMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
