package container

import (
	"log"

	"github.com/rivo/tview"
)

var containerMap = make(map[string]interface{})
var App *tview.Application

func init() {
	containerMap["afterInitialize"] = []func(){}
}

func AddComponent(name string, component interface{}) {
	containerMap[name] = component
}

func GetComponent(name string) (interface{}, bool) {
	v, ok := containerMap[name]
	return v, ok
}

func AfterInitialize(function func()) {
	containerMap["afterInitialize"] = append(containerMap["afterInitialize"].([]func()), function)
}

func ExecuteAfterInitialize() {
	functions := containerMap["afterInitialize"].([]func())

	for i := range functions {
		log.Println("executing")
		functions[i]()
	}
}
