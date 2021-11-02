package keylist

import (
	"riakg/components/shared"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var keysFontColor = tcell.NewRGBColor(200, 200, 200)

func NewKeyList() *tview.List {
	keyList := tview.NewList()
	shared.SetBaseStyle(keyList, "Keys")
	keyList.ShowSecondaryText(false)
	keyList.SetMainTextColor(keysFontColor)

	return keyList
}
