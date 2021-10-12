package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"riakg/riakapi"
)

var backgroundColor = tcell.NewRGBColor(26, 27, 38)
var bucketsFontColor = tcell.NewRGBColor(47, 196, 222)
var borderColor = tcell.NewRGBColor(59, 66, 97)

//var keysFontColor = tcell.NewRGBColor(187, 154, 247)
var keysFontColor = tcell.NewRGBColor(200, 200, 200)

func empty() {}

func setSelectedFunct(app *tview.Application, tree *tview.TreeView, keyList *tview.List) {
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		keyList.Clear()
		keys := riakapi.GetBucketKeys(node.GetText())

		for i := range keys {
			keyList.AddItem(keys[i], "", 0, empty)
		}
		app.SetFocus(keyList)
	})
}

func setSelectedKeyHandler(app *tview.Application, keyList *tview.List, keyValue *tview.TextView) {
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

	// Tree declaracion
	bucketTree := tview.NewTreeView()
	bucketTree.SetBorder(true)
	bucketTree.SetBackgroundColor(backgroundColor)
	bucketTree.SetBorderColor(borderColor)
	bucketTree.SetTitle(riakapi.Host + ":" + riakapi.Port)

	// Key list declaration
	keyList := tview.NewList()
	keyList.ShowSecondaryText(false)
	keyList.SetBackgroundColor(backgroundColor)
	keyList.SetBorder(true).SetTitle("Keys")
	keyList.SetBorderColor(borderColor)
	keyList.SetMainTextColor(keysFontColor)
	keyList.SetDoneFunc(func() {
		app.SetFocus(bucketTree)
	})

	// Key Value declaration
	keyView := tview.NewTextView().SetWrap(false)
	keyView.SetBorder(true).SetTitle("Value")
	keyView.SetBackgroundColor(backgroundColor)
	keyView.SetBorderColor(borderColor)
	keyView.SetDynamicColors(true).SetRegions(true)
	keyView.SetScrollable(true)
	keyView.SetWrap(true)

	flex.AddItem(bucketTree, 0, 1, true)
	flex.AddItem(keyList, 0, 1, false)
	flex.AddItem(keyView, 0, 2, false)

	// Set bindings
	setSelectedFunct(app, bucketTree, keyList)

	keyList.SetSelectedFunc(func(idx int, key, secondary string, shortcut rune) {
		currentBucket := bucketTree.GetCurrentNode().GetText()
		value := riakapi.GetKeyValue(currentBucket, key)
		keyView.SetText(value)
		app.SetFocus(keyView)
	})
	keyView.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(keyList)
	})

	fillBuckets(bucketTree)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func fillBuckets(bucketTree *tview.TreeView) {
	buckets := riakapi.GetBuckets()

	rootDir := "Buckets"
	root := tview.NewTreeNode(rootDir).SetColor(bucketsFontColor)
	bucketTree.SetRoot(root).SetCurrentNode(root)

	for v := range buckets.Bukckets {
		root.AddChild(tview.NewTreeNode(buckets.Bukckets[v]).SetColor(bucketsFontColor))
	}
}
