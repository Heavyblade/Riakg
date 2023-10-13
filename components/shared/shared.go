package shared

import (
	"riakg/components/container"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var BackgroundColor = tcell.ColorDefault
var BucketsFontColor = tcell.NewRGBColor(47, 196, 222)
var BorderColor = tcell.NewRGBColor(59, 66, 97)
var SelectedColor = tcell.NewRGBColor(251, 158, 101)
var FontBlueColor = tcell.NewRGBColor(118, 157, 239)

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
	component.SetBackgroundColor(BackgroundColor)

	if len(title) > 0 {
		component.SetTitle(title)
		component.SetBorder(true)
		component.SetBorderColor(BorderColor)
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
	source.(BaseSettabler).SetBorderColor(BorderColor)
	destination.(BaseSettabler).SetBorderColor(SelectedColor)
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

func WrappInPages(pageList map[string]tview.Primitive) *tview.Pages {
	pages := tview.NewPages()
	justFirst := true

	for keym, page := range pageList {
		pages.AddPage(keym, page, false, justFirst)
		justFirst = false
	}

	return pages
}

func ConfirmAction(message string, origin tview.Primitive, yes, no func(modal *tview.Modal)) *tview.Modal {
	untypedFlex, _ := container.GetComponent("mainLayout")
	flex := untypedFlex.(*tview.Flex)

	modal := tview.NewModal()
	modal.SetText(message).
		AddButtons([]string{"Yes", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			flex.RemoveItem(modal)
			if buttonLabel == "Yes" && yes != nil {
				yes(modal)
			}
			if buttonLabel == "Cancel" && no != nil {
				no(modal)
			}
			container.App.SetFocus(origin)
		})
	modal.SetBackgroundColor(BackgroundColor)
	modal.SetButtonTextColor(FontBlueColor)

	flex.AddItem(modal, 2, 1, true)
	container.App.SetFocus(modal)

	return modal
}
