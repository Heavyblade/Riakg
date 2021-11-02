package buckettree

import (
	"riakg/components/shared"
	"riakg/riakapi"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tree *tview.TreeView
var bucketsFontColor = tcell.NewRGBColor(47, 196, 222)

func NewBucketTree() *tview.TreeView {
	bucketTree := tview.NewTreeView()
	shared.SetBaseStyle(bucketTree, riakapi.Host+":"+riakapi.Port)

	tree = bucketTree
	return tree
}

func FillBuckets() {
	buckets := riakapi.GetBuckets()

	rootDir := "Buckets"
	root := tview.NewTreeNode(rootDir).SetColor(bucketsFontColor)
	tree.SetRoot(root).SetCurrentNode(root)

	for v := range buckets.Bukckets {
		root.AddChild(tview.NewTreeNode(buckets.Bukckets[v]).SetColor(bucketsFontColor))
	}
}
