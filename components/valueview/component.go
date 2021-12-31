package valueview

import (
	"fmt"
	"riakg/components/container"
	"riakg/components/shared"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tidwall/pretty"
	"golang.design/x/clipboard"
)

func init() {
	component := NewValueView()
	container.AddComponent("valueView", component)

	wrapped := shared.WrapWithShortCuts(component, []string{"Ctrl-y: Copy value", "Ctrl-s: Change key"})
	container.AddComponent("WrappedvalueView", wrapped)

	container.AfterInitialize(func() {
		destination, _ := container.GetComponent("bucketTree")

		shared.SetTabDestination(component, destination.(*tview.TreeView))
		fun := component.GetInputCapture()

		component.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			fun(event)
			if event.Key() == tcell.KeyCtrlY {
				clipboard.Write(clipboard.FmtText, []byte(component.GetText(true)))
			}
			if event.Key() == tcell.KeyCtrlV {
				prettified := pretty.Pretty(clipboard.Read(clipboard.FmtText))
				highlighted := pretty.Color([]byte(prettified), nil)

				component.Clear()
				w := tview.ANSIWriter(component)
				fmt.Fprint(w, string(highlighted))
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
