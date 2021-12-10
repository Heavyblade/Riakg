package shared

import (
	"riakg/components/container"

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

type InputCapturabler interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
}

func SetBaseStyle(component BaseSettabler, title string) {
	component.SetBorder(true)
	component.SetBackgroundColor(backgroundColor)
	component.SetBorderColor(borderColor)
	component.SetTitle(title)
}

func SetTabDestination(source InputCapturabler, destination tview.Primitive) {
	source.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			container.App.Stop()
		case tcell.KeyTAB:
			container.App.SetFocus(destination)
		}

		return event
	})
}
