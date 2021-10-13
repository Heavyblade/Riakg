package main

import (
	"fmt"
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

func setSelectedBucketHandler(app *tview.Application, tree *tview.TreeView, keyList *tview.List) {
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		keyList.Clear()
		keys := riakapi.GetBucketKeys(node.GetText())

		for i := range keys {
			keyList.AddItem(keys[i], "", 0, empty)
		}
		app.SetFocus(keyList)
	})
}

func setSelectedKeyHandler(app *tview.Application, bucketTree *tview.TreeView, keyList *tview.List, valueView *tview.TextView) {
	keyList.SetSelectedFunc(func(idx int, key, secondary string, shortcut rune) {
		currentBucket := bucketTree.GetCurrentNode().GetText()
		value := riakapi.GetKeyValue(currentBucket, key)
		valueView.Clear()
		w := tview.ANSIWriter(valueView)
		fmt.Fprint(w, value)
		app.SetFocus(valueView)
	})
}

type BaseSettabler interface {
	SetBorder(bool) *tview.Box
	SetBackgroundColor(tcell.Color) *tview.Box
	SetBorderColor(tcell.Color) *tview.Box
	SetTitle(string) *tview.Box
}

func setBaseStyle(component BaseSettabler, title string) {
	component.SetBorder(true)
	component.SetBackgroundColor(backgroundColor)
	component.SetBorderColor(borderColor)
	component.SetTitle(title)
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
	setBaseStyle(bucketTree, riakapi.Host+":"+riakapi.Port)

	// Key list declaration
	keyList := tview.NewList()
	setBaseStyle(keyList, "Keys")
	keyList.ShowSecondaryText(false)
	keyList.SetMainTextColor(keysFontColor)
	keyList.SetDoneFunc(func() {
		app.SetFocus(bucketTree)
	})

	// Key Value declaration
	valueView := tview.NewTextView().SetWrap(false)
	setBaseStyle(valueView, "Keys")
	valueView.SetDynamicColors(true)
	valueView.SetScrollable(true)
	valueView.SetWrap(false)

	flex.AddItem(bucketTree, 0, 1, true)
	flex.AddItem(keyList, 0, 1, false)
	flex.AddItem(valueView, 0, 2, false)

	// Set bindings
	setSelectedBucketHandler(app, bucketTree, keyList)
	setSelectedKeyHandler(app, bucketTree, keyList, valueView)

	valueView.SetDoneFunc(func(key tcell.Key) {
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
