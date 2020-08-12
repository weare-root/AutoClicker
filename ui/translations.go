package ui

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gotk3/gotk3/glib"
	"github.com/weare-root/AutoClicker/errs"
)

func readTranslation(lang string) map[string]interface{} {
	file, err := ioutil.ReadFile(getPath("resources", "translations", lang+".json"))
	errs.ErrorCheck(err)

	data := map[string]interface{}{}

	err = json.Unmarshal(file, &data)
	errs.ErrorCheck(err)

	return data
}

func applyTranslation(data map[string]interface{}) {
	glib.IdleAdd(func() {

	})
}
