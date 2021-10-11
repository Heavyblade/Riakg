package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"riakg/riakapi"
)

func empty() {}

func setSelectedFunct(app *tview.Application, tree *tview.TreeView, keyList *tview.List) {

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		keyList.Clear()
		keys := riakapi.GetBucketKeys(node.GetText())

		for i := range keys {
			keyList.AddItem(keys[i], "", 0, empty)
		}

		app.SetFocus(keyList)
		//children := node.GetChildren()
		//if len(children) == 0 {
		//// Load and show files in this directory.
		//path := reference.(string)
		//add(node, path)
		//} else {
		//// Collapse if visible, expand if collapsed.
		//node.SetExpanded(!node.IsExpanded())
		//}
	})
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
	tree := tview.NewTreeView()
	tree.SetBorder(true)
	tree.SetBackgroundColor(tcell.NewRGBColor(20, 20, 20))
	tree.SetBorderColor(tcell.NewRGBColor(40, 40, 40))

	// Root element
	rootDir := "Buckets"
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	tree.SetRoot(root).SetCurrentNode(root)

	buckets := riakapi.GetBuckets()

	for v := range buckets.Bukckets {
		root.AddChild(tview.NewTreeNode(buckets.Bukckets[v]))
	}

	// Key list
	keyList := tview.NewList()
	keyList.ShowSecondaryText(false)
	keyList.SetBorder(true).SetTitle("Keys")
	keyList.SetDoneFunc(func() {
		app.SetFocus(tree)
		//keyList.Clear()
		//app.SetFocus(databases)
	})

	setSelectedFunct(app, tree, keyList)

	flex.AddItem(tree, 0, 1, true)
	flex.AddItem(keyList, 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
