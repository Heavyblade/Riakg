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

	container.AfterInitialize(func() {
		destination, _ := container.GetComponent("bucketTree")

		shared.SetTabDestination(component, destination.(*tview.TreeView))
		fun := component.GetInputCapture()

		component.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			fun(event)
			if event.Key() == tcell.KeyCtrlY {
				clipboard.Write(clipboard.FmtText, []byte(component.GetText(true)))
			}
			return event
		})
	})
}

func NewValueView() *tview.TextView {
	valueView := tview.NewTextView().SetWrap(false)
	shared.SetBaseStyle(valueView, "Value")
	valueView.SetDynamicColors(true)
	valueView.SetScrollable(true)
	valueView.SetWrap(false)

	return valueView
}
