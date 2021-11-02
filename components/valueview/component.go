package valueview

import (
	"riakg/components/shared"

	"github.com/rivo/tview"
)

func NewValueView() *tview.TextView {
	valueView := tview.NewTextView().SetWrap(false)
	shared.SetBaseStyle(valueView, "Value")
	valueView.SetDynamicColors(true)
	valueView.SetScrollable(true)
	valueView.SetWrap(false)

	return valueView
}
