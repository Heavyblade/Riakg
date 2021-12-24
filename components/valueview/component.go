package valueview

import (
	"riakg/components/container"
	"riakg/components/shared"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

func init() {
	component := NewValueView()
	container.AddComponent("valueView", component)

	wrapped := shared.WrapWithShortCuts(component, []string{"Ctrl-y: Copy value"})
	container.AddComponent("WrappedvalueView", wrapped)

	container.AfterInitialize(func() {
		destination, _ := container.GetComponent("bucketTree")

		shared.SetTabDestination(component, destination.(*tview.TreeView))
		fun := component.GetInputCapture()

		component.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			fun(event)
			if event.Key() == tcell.KeyCtrlY {
				clipboard.Write(clipboard.FmtText, []byte(component.GetText(true)))
			}
			if event.Key() == tcell.KeyCtrlS {
				container.App.SetRoot(changeFunction(), true)
			}
			return event
		})
	})
}

func changeFunction() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			container.App.Stop()
		})

	return form
}

func NewValueView() *tview.TextView {
	valueView := tview.NewTextView().SetWrap(false)
	shared.SetBaseStyle(valueView, "Value")
	valueView.SetDynamicColors(true)
	valueView.SetScrollable(true)
	valueView.SetWrap(false)

	return valueView
}
