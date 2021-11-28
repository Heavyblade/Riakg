package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"riakg/components/buckettree"
	"riakg/components/keylist"
	"riakg/components/valueview"
	"riakg/riakapi"
)

type InputCapturabler interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
}

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex()

	bucketTree := buckettree.NewBucketTree()
	keyList := keylist.NewKeyList()
	valueView := valueview.NewValueView()

	flex.AddItem(bucketTree, 0, 1, true)
	flex.AddItem(keyList, 0, 1, false)
	flex.AddItem(valueView, 0, 2, false)

	// Set bindings
	setSelectedBucketHandler(app, bucketTree, keyList)
	setSelectedKeyHandler(app, bucketTree, keyList, valueView)
	setTabDestination(app, bucketTree, keyList)
	setTabDestination(app, keyList, valueView)
	setTabDestination(app, valueView, bucketTree)

	buckettree.FillBuckets()

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func empty() {}

func setSelectedBucketHandler(app *tview.Application, tree *tview.TreeView, keyList *tview.List) {
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		keyList.Clear()
		keys := riakapi.GetBucketKeys(node.GetText())

		for i := range keys {
			keyList.AddItem(keys[i], node.GetText(), 0, empty)
		}
		app.SetFocus(keyList)
	})
}

func setSelectedKeyHandler(app *tview.Application, bucketTree *tview.TreeView, keyList *tview.List, valueView *tview.TextView) {
	keyList.SetChangedFunc(func(idx int, key, secondary string, shortcut rune) {
		currentBucket := bucketTree.GetCurrentNode().GetText()
		value := riakapi.GetKeyValue(currentBucket, key)
		valueView.Clear()
		w := tview.ANSIWriter(valueView)
		fmt.Fprint(w, value)
	})
}

func setTabDestination(app *tview.Application, source InputCapturabler, destination tview.Primitive) {
	source.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if w, ok := source.(*tview.List); ok && event.Key() == tcell.KeyCtrlD {
			idx := w.GetCurrentItem()
			key, bucket := w.GetItemText(idx)

			if riakapi.DeleteKey(bucket, key) {
				// Needed due a bug on the RemoveItem function when the item to remove
				// is the first one on the list
				if idx == 0 {
					w.SetCurrentItem(1)
				}
				w.RemoveItem(idx)
			}
			return event
		}
		if event.Key() == tcell.KeyTAB {
			app.SetFocus(destination)
		}
		return event
	})
}
