package keylist

import (
	"fmt"
	"riakg/components/container"
	"riakg/components/shared"
	"riakg/riakapi"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func init() {
	component := NewKeyList()
	container.AddComponent("keyList", component)

	wrapped := shared.WrapWithShortCuts(component, []string{"Ctrl-y: Copy key", "Ctrl-d: Delete key", "Ctrl-x: Delete all keys"})
	container.AddComponent("WrappedkeyList", wrapped)

	container.AfterInitialize(func() {
		valueViewUntyped, _ := container.GetComponent("valueView")
		bucketTreeUntyped, _ := container.GetComponent("bucketTree")

		valueView := valueViewUntyped.(*tview.TextView)
		bucketTree := bucketTreeUntyped.(*tview.TreeView)

		shared.SetTabDestination(component, valueView)
		tabCapturefunc := component.GetInputCapture()

		component.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			tabCapturefunc(event)

			switch event.Key() {
			case tcell.KeyCtrlD:
				shared.ConfirmAction("Are you sure?", component, func(modal *tview.Modal) {
					idx := component.GetCurrentItem()
					key, bucket := component.GetItemText(idx)

					if riakapi.DeleteKey(bucket, key) {
						// Needed due a bug on the RemoveItem function when the item to remove
						// is the first one on the component
						if idx == 0 {
							component.SetCurrentItem(1)
						}
						component.RemoveItem(idx)
					}
					valueView.Clear()
				}, nil)
			case tcell.KeyCtrlX:
				shared.ConfirmAction("Are you sure?", component, func(modal *tview.Modal) {
					itemCount := component.GetItemCount()

					for idx := 0; idx < itemCount; idx++ {
						key, bucket := component.GetItemText(0)

						if riakapi.DeleteKey(bucket, key) {
							component.SetCurrentItem(1)
							component.RemoveItem(0)
						}
					}
					valueView.Clear()
				}, nil)
			case tcell.KeyCtrlY:
				key, _ := component.GetItemText(component.GetCurrentItem())
				clipboard.WriteAll(key)
			}

			// This prevents that tab bubbles to the main goes to the main handler and
			// moves the cursor to the next item
			if event.Key() == tcell.KeyTAB {
				return nil
			}
			return event
		})

		component.SetChangedFunc(func(idx int, key, secondary string, shortcut rune) {
			currentBucket := bucketTree.GetCurrentNode().GetText()
			value := riakapi.GetKeyValue(currentBucket, key)
			valueView.Clear()
			valueView.ScrollToBeginning()
			w := tview.ANSIWriter(valueView)
			fmt.Fprint(w, value)
		})
	})
}

var keysFontColor = tcell.NewRGBColor(200, 200, 200)

func NewKeyList() *tview.List {
	keyList := tview.NewList()
	shared.SetBaseStyle(keyList, "Keys")
	keyList.ShowSecondaryText(false)
	keyList.SetMainTextColor(keysFontColor)

	return keyList
}
