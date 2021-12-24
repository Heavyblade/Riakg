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
	keyListUntyped, _ := container.GetComponent("WrappedkeyList")
	valueViewUntyped, _ := container.GetComponent("WrappedvalueView")

	container.ExecuteAfterInitialize()

	flex := tview.NewFlex()
	flex.AddItem(bucketTreeUntyped.(*tview.TreeView), 0, 1, true)
	flex.AddItem(keyListUntyped.(*tview.Flex), 0, 1, false)
	flex.AddItem(valueViewUntyped.(*tview.Pages), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
