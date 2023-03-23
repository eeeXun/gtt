package ui

import (
	"github.com/rivo/tview"
)

type UICycle struct {
	widget []tview.Primitive
	index  int8
	len    int8
}

func NewUICycle(widgets ...tview.Primitive) *UICycle {
	var w []tview.Primitive

	for _, widget := range widgets {
		w = append(w, widget)
	}

	return &UICycle{
		widget: w,
		index:  0,
		len:    int8(len(w)),
	}
}

func (ui *UICycle) Increase() {
	ui.index = (ui.index + 1) % ui.len
}

func (ui *UICycle) Decrease() {
	ui.index = ((ui.index-1)%ui.len + ui.len) % ui.len
}

func (ui *UICycle) GetCurrentUI() tview.Primitive {
	return ui.widget[ui.index]
}
