package shared

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var backgroundColor = tcell.NewRGBColor(26, 27, 38)
var bucketsFontColor = tcell.NewRGBColor(47, 196, 222)
var borderColor = tcell.NewRGBColor(59, 66, 97)

type BaseSettabler interface {
	SetBorder(bool) *tview.Box
	SetBackgroundColor(tcell.Color) *tview.Box
	SetBorderColor(tcell.Color) *tview.Box
	SetTitle(string) *tview.Box
}

func SetBaseStyle(component BaseSettabler, title string) {
	component.SetBorder(true)
	component.SetBackgroundColor(backgroundColor)
	component.SetBorderColor(borderColor)
	component.SetTitle(title)
}
