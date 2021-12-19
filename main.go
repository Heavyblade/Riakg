package main

import (
	"log"
	"os"

	"github.com/rivo/tview"

	_ "riakg/components/buckettree"
	"riakg/components/container"
	_ "riakg/components/keylist"
	"riakg/components/shared"
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
	keyListUntyped, _ := container.GetComponent("keyList")
	valueViewUntyped, _ := container.GetComponent("valueView")

	container.ExecuteAfterInitialize()

	flex := tview.NewFlex()
	flex.AddItem(bucketTreeUntyped.(*tview.TreeView), 0, 1, true)
	flex.AddItem(wrappWithOptions(keyListUntyped.(*tview.List), "Ctrl+d Delete key"), 0, 1, false)
	flex.AddItem(wrappWithOptions(valueViewUntyped.(*tview.TextView), "Crrl+y Copy content"), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func wrappWithOptions(comp tview.Primitive, helpText string) tview.Primitive {
	text := tview.NewTextView().SetText(helpText)
	text.SetTextAlign(tview.AlignCenter)

	shared.SetBaseStyle(text, "")
	flex := tview.NewFlex()
	flex.AddItem(comp, 0, 1, true)
	flex.AddItem(text, 2, 1, false)
	flex.SetDirection(tview.FlexRow)

	return flex
}
