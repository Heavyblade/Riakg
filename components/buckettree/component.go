package buckettree

import (
	"riakg/components/container"
	"riakg/components/shared"
	"riakg/riakapi"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var bucketsFontColor = tcell.NewRGBColor(47, 196, 222)

func init() {
	component := NewBucketTree()
	container.AddComponent("bucketTree", component)

	container.AfterInitialize(func() {
		keyListUntyped, _ := container.GetComponent("keyList")
		keyList := keyListUntyped.(*tview.List)

		shared.SetTabDestination(component, keyList)
		component.SetSelectedFunc(selectedLeafHandler(keyList, component))

		fillBuckets()
	})
}

func selectedLeafHandler(keyList *tview.List, bucketTree *tview.TreeView) func(node *tview.TreeNode) {
	return func(node *tview.TreeNode) {
		valueViewUntyped, _ := container.GetComponent("valueView")
		valueView := valueViewUntyped.(*tview.TextView)

		keyList.Clear()
		valueView.Clear()
		keys := riakapi.GetBucketKeys(node.GetText())

		for i := range keys {
			keyList.AddItem(keys[i], node.GetText(), 0, func() {})
		}
		container.App.SetFocus(keyList)
	}
}

func NewBucketTree() *tview.TreeView {
	bucketTree := tview.NewTreeView()
	shared.SetBaseStyle(bucketTree, riakapi.Host+":"+riakapi.Port)

	return bucketTree
}

func fillBuckets() {
	buckets := riakapi.GetBuckets()
	bucketTreeUntyped, _ := container.GetComponent("bucketTree")
	tree := bucketTreeUntyped.(*tview.TreeView)

	rootDir := "Buckets"
	root := tview.NewTreeNode(rootDir).SetColor(bucketsFontColor)
	tree.SetRoot(root).SetCurrentNode(root)

	for v := range buckets.Bukckets {
		root.AddChild(tview.NewTreeNode(buckets.Bukckets[v]).SetColor(bucketsFontColor))
	}
}
