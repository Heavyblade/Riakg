package main

import (
	"log"
	"os"

	"github.com/rivo/tview"

	_ "riakg/components/buckettree"
	"riakg/components/container"
	_ "riakg/components/keylist"
	_ "riakg/components/valueview"
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func main() {
	app := tview.NewApplication()
	container.App = app

	bucketTreeUntyped, _ := container.GetComponent("bucketTree")
	bucketTree := bucketTreeUntyped.(*tview.TreeView)

	keyListUntyped, _ := container.GetComponent("keyList")
	keyList := keyListUntyped.(*tview.List)

	valueViewUntyped, _ := container.GetComponent("valueView")
	valueView := valueViewUntyped.(*tview.TextView)

	container.ExecuteAfterInitialize()

	flex := tview.NewFlex()
	flex.AddItem(bucketTree, 0, 1, true)
	flex.AddItem(keyList, 0, 1, false)
	flex.AddItem(valueView, 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
