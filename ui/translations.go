package ui

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gotk3/gotk3/glib"
	"github.com/weare-root/AutoClicker/errs"
)

func readTranslation(lang string) map[string]string {
	file, err := ioutil.ReadFile(getPath("resources", "translations", lang+".json"))
	errs.ErrorCheck(err)

	data := make(map[string]string)

	err = json.Unmarshal(file, &data)
	errs.ErrorCheck(err)

	return data
}

func applyTranslation(data map[string]string) {
	glib.IdleAdd(func() {
		lbActivation.SetText(data["activation"])
		lbActivationKey.SetText(data["activationkey"])
		btKey.SetLabel(data["choosekey"])
		lbActivationMode.SetText(data["activationmode"])
		rbHold.SetLabel(data["hold"])
		rbSwitch.SetLabel(data["switch"])
		lbTimespan.SetText(data["timespan"])
		lbRatio.SetText(data["ratio"])
		lbCPSLower.SetText(data["cpslower"])
		lbCPSHigher.SetText(data["cpshigher"])
		lbClicks.SetText(data["clicks"])
		rbLeft.SetLabel(data["leftmb"])
		rbMiddle.SetLabel(data["middlemb"])
		rbRight.SetLabel(data["rightmb"])
		rbCustom.SetLabel(data["custom"])
		btCustom.SetLabel(data["choosekey"])
		lbLanguage.SetText(data["language"])
		lbChooseLanguage.SetText(data["chooselanguage"])
	})
}
