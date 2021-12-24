package valueview

import (
	"riakg/components/container"
	"riakg/components/shared"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

func init() {
	component := NewValueView()
	container.AddComponent("valueView", component)

	wrapped := shared.WrapWithShortCuts(component, []string{"Ctrl-y: Copy value", "Ctrl-s: Change key"})
	pages := tview.NewPages()
	modal := buildModal(pages)
	pages.AddPage("valueView", wrapped, true, true)
	pages.AddPage("modal", modal, true, false)

	container.AddComponent("WrappedvalueView", pages)
	container.AfterInitialize(func() {
		destination, _ := container.GetComponent("bucketTree")

		shared.SetTabDestination(component, destination.(*tview.TreeView))
		fun := component.GetInputCapture()

		component.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			fun(event)
			if event.Key() == tcell.KeyCtrlY {
				clipboard.Write(clipboard.FmtText, []byte(component.GetText(true)))
			}
			if event.Key() == tcell.KeyCtrlS {
				pages.SwitchToPage("modal")
				container.App.SetFocus(modal)
			}
			return event
		})
	})
}

func returnToValueView(pages *tview.Pages) func() {
	destination, _ := container.GetComponent("valueView")
	return func() {
		pages.SwitchToPage("valueView")
		container.App.SetFocus(destination.(*tview.TextView))
	}
}

func buildModal(pages *tview.Pages) *tview.Form {
	form := tview.NewForm().
		AddInputField("Key", "", 50, nil, nil).
		AddInputField("Value", "", 50, nil, nil).
		AddButton("Save", returnToValueView(pages)).
		AddButton("Quit", returnToValueView(pages))

	shared.SetBaseStyle(form, "Update Key")

	return form
}

func NewValueView() *tview.TextView {
	valueView := tview.NewTextView().SetWrap(false)
	shared.SetBaseStyle(valueView, "Value")
	valueView.SetDynamicColors(true)
	valueView.SetScrollable(true)
	valueView.SetWrap(false)

	return valueView
}
