package shared

import (
	"riakg/components/container"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var backgroundColor = tcell.NewRGBColor(26, 27, 38)
var bucketsFontColor = tcell.NewRGBColor(47, 196, 222)
var borderColor = tcell.NewRGBColor(59, 66, 97)
var selectedColor = tcell.NewRGBColor(251, 158, 101)

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
	component.SetBackgroundColor(backgroundColor)

	if len(title) > 0 {
		component.SetTitle(title)
		component.SetBorder(true)
		component.SetBorderColor(borderColor)
	}
}

func SetTabDestination(source InputCapturabler, destination tview.Primitive) {
	source.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			container.App.Stop()
		case tcell.KeyTAB:
			SetTabFocusAndBorders(source, destination)
		}

		return event
	})
}

func SetTabFocusAndBorders(source InputCapturabler, destination tview.Primitive) {
	container.App.SetFocus(destination)
	source.(BaseSettabler).SetBorderColor(borderColor)
	destination.(BaseSettabler).SetBorderColor(selectedColor)
}

func WrapWithShortCuts(comp tview.Primitive, helpText []string) *tview.Flex {
	text := tview.NewTextView().SetText(strings.Join(helpText, "\n"))
	text.SetTextAlign(tview.AlignLeft)

	text.SetBorderPadding(0, 0, 1, 0)
	SetBaseStyle(text, "")
	flex := tview.NewFlex()
	flex.AddItem(comp, 0, 1, true)
	flex.AddItem(text, 1+len(helpText), 1, false)
	flex.SetDirection(tview.FlexRow)

	return flex
}

func WrappInPages(pageList map[string]tview.Primitive) tview.Primitive {
	pages := tview.NewPages()
	justFirst := true

	for keym, page := range pageList {
		pages.AddPage(keym, page, false, justFirst)
		justFirst = false
	}

	return pages
}
