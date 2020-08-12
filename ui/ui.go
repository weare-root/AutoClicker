package ui

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	hook "github.com/robotn/gohook"
	"github.com/weare-root/AutoClicker/autoclicker"
	"github.com/weare-root/AutoClicker/errs"
)

var gBuilder *gtk.Builder

// variables for runtime
var (
	activationBtn = ""
	listening     = false
	listeningBtn  *gtk.Button
	listeningEn   *gtk.Entry
	shouldListen  = false
)

// definitions for the widgets
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

	cbLanguage *gtk.ComboBoxText

	lbActivation     *gtk.Label
	lbActivationKey  *gtk.Label
	lbActivationmode *gtk.Label
	lbTimespan       *gtk.Label
	lbRatio          *gtk.Label
	lbCPSLower       *gtk.Label
	lbCPSHigher      *gtk.Label
	lbClicks         *gtk.Label
	lbLanguage       *gtk.Label
)

// OnActivate is used for when the application starts up
func OnActivate(application *gtk.Application) {
	log.Println("UI is being built")
	// get the builder
	builder, err := gtk.BuilderNewFromFile(getPath("resources", "ui", "main.glade"))
	errs.ErrorCheck(err)
	gBuilder = builder

	// add signals
	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"activation_btn_clicked": activationBtnClicked,
		"custom_btn_clicked":     customBtnClicked,
		"rbCustom_toggled":       rbCustomToggled,
		"lbLanguage_changed":     lbLanguageChanged,
	}
	builder.ConnectSignals(signals)
	log.Println("connected signals")

	// get the window
	win := getApplicationWindow("main_window")

	err = win.SetIconFromFile(getPath("resources", "ui", "icon.png"))
	errs.LogOnError(err, "could not set icon")

	//add styling
	screen := win.GetScreen()

	// add own customizations (not much lol)
	provider, err := gtk.CssProviderNew()
	errs.ErrorCheck(err)
	err = provider.LoadFromPath(getPath("resources", "ui", "main.css"))
	errs.ErrorCheck(err)
	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	log.Println("Applied custom styles")

	// i think this is so the window can listen to keyboard events
	win.AddEvents(int(gdk.KEY_PRESS_MASK))

	// add the window to the application
	application.AddWindow(win)
	// Show the Window and all of its components.
	win.Show()
	log.Println("built the UI and showing it")

	getWidgets()

	go listenToKeyboard()

	// go startClicker(10, 15, 0.5, 1000, "left")
	// time.Sleep(time.Second * 3)
	// stopClicker()
}

// getWidgets gets all the necessary widgets for later use
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

	log.Println("got all necessary widgets")
}

// listenToKeyboard handles the logic for listening to the keyboard
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
				// check if the button for the listening key is pressed
				if listeningBtn == btKey {
					activationBtn = raw2key[ev.Rawcode]
				}
				execMainThread(func() {
					listeningEn.SetText(raw2key[ev.Rawcode])
				})
			}
			continue
		}

		if ev.Rawcode == keytoraw[activationBtn] && activationBtn != "" {
			// check if its a hold for the holding mode
			if ev.Kind == hook.KeyHold && isHoldingMode() && !isDown {
				// start the holding if it isnt already started
				log.Println("HOLDING start")
				isDown = true
				go autoclicker.StartClicker(float32(adLower.GetValue()), float32(adHigher.GetValue()), float32(adRatio.GetValue()/100.0), int(adTimespan.GetValue()), getKey())
			} else if ev.Kind == hook.KeyUp {
				if getCurrentMillis()-lastChanged < 500 {
					continue
				}
				lastChanged = getCurrentMillis()

				if isHoldingMode() {
					log.Println("HOLDING stop")
					// stop when holding mode is on
					isDown = false
					autoclicker.StopClicker()
				} else if !isDown {
					log.Println("SWITCHING start")
					// start when switching mode is on
					isDown = true
					go autoclicker.StartClicker(float32(adLower.GetValue()), float32(adHigher.GetValue()), float32(adRatio.GetValue()/100.0), int(adTimespan.GetValue()), getKey())
				} else {
					// stop when switching mode is on
					log.Println("SWITCHING stop")
					isDown = false
					autoclicker.StopClicker()
				}
			}

		}
	}
}

func getCurrentMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
