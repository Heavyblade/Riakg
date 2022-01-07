package main

import (
	"flag"
	"log"
	"os"

	"github.com/rivo/tview"

	_ "riakg/components/buckettree"
	"riakg/components/container"
	_ "riakg/components/keylist"
	_ "riakg/components/valueview"
	"riakg/riakapi"
)

func init() {
	getParams()
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func main() {
	app := tview.NewApplication()
	container.App = app
	flex := tview.NewFlex()
	container.AddComponent("mainLayout", flex)

	bucketTreeUntyped, _ := container.GetComponent("bucketTree")
	keyListUntyped, _ := container.GetComponent("WrappedkeyList")
	valueViewUntyped, _ := container.GetComponent("WrappedvalueView")

	container.ExecuteAfterInitialize()

	flex.AddItem(bucketTreeUntyped.(*tview.TreeView), 0, 1, true)
	flex.AddItem(keyListUntyped.(*tview.Flex), 0, 1, false)
	flex.AddItem(valueViewUntyped.(*tview.Flex), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func getParams() {
	host := flag.String("host", "localhost", "server ip or domain")
	port := flag.String("port", "8098", "server port")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "password")

	flag.Parse()

	riakapi.Host = *host
	riakapi.Port = *port
	riakapi.Username = *username
	riakapi.Password = *password
}
