package valueview

import (
	"fmt"
	"riakg/components/container"
	"riakg/components/shared"
	"riakg/riakapi"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tidwall/pretty"
)

func init() {
	component := NewValueView()
	container.AddComponent("valueView", component)

	wrapped := shared.WrapWithShortCuts(component, []string{"Ctrl-y: Copy", "Ctrl-s: Save", "Ctrl-v: Paste new value"})
	container.AddComponent("WrappedvalueView", wrapped)

	container.AfterInitialize(func() {
		destination, _ := container.GetComponent("bucketTree")
		bucketTree := destination.(*tview.TreeView)

		shared.SetTabDestination(component, bucketTree)
		fun := component.GetInputCapture()

		component.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			fun(event)

			switch event.Key() {
			case tcell.KeyCtrlY:
				copyValueToClipboard(component)
			case tcell.KeyCtrlV:
				value, _ := clipboard.ReadAll()
				updateValue(component, []byte(value))
			case tcell.KeyCtrlS:
				shared.ConfirmAction("Are you sure?", component, func(modal *tview.Modal) {
					keyListUntyped, _ := container.GetComponent("keyList")
					keyList := keyListUntyped.(*tview.List)

					currentBucket := bucketTree.GetCurrentNode().GetText()
					key, _ := keyList.GetItemText(keyList.GetCurrentItem())
					riakapi.UpdateKeyValue(currentBucket, key, component.GetText(true))
				}, nil)
			}
			return event
		})
	})
}

func copyValueToClipboard(component *tview.TextView) {
	clipboard.WriteAll(component.GetText(true))
}

func updateValue(component *tview.TextView, newValue []byte) {
	prettified := pretty.Pretty(newValue)
	highlighted := pretty.Color([]byte(prettified), nil)
	component.Clear()
	w := tview.ANSIWriter(component)
	fmt.Fprint(w, string(highlighted))
}

func NewValueView() *tview.TextView {
	valueView := tview.NewTextView().SetWrap(false)
	shared.SetBaseStyle(valueView, "Value")
	valueView.SetDynamicColors(true)
	valueView.SetScrollable(true)
	valueView.SetWrap(false)

	return valueView
}
