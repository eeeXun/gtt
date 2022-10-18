package main

import (
	"github.com/gdamore/tcell/v2"
)

func ui_init() {
	src_box.SetBorder(true).
		SetTitle(translator.src_lang).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color).
		SetBackgroundColor(window.src.background_color)
	src_box.SetTextStyle(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color))
	src_box.SetSelectedStyle(tcell.StyleDefault.
		Background(window.src.selected_color).
		Foreground(window.src.foreground_color))
	src_box.SetInputCapture(InputHandle)

	dest_box.SetBorder(true).
		SetTitle(translator.dest_lang).
		SetBorderColor(window.dest.border_color).
		SetTitleColor(window.dest.border_color).
		SetBackgroundColor(window.dest.background_color)
	dest_box.SetTextColor(window.dest.foreground_color)
}

func InputHandle(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	switch key {
	case tcell.KeyCtrlJ:
		result, err := translator.Translate(src_box.GetText())
		if err != nil {
			dest_box.SetText(err.Error())
		} else {
			dest_box.SetText(result)
		}
	case tcell.KeyCtrlQ:
		src_box.SetText("", true)
	case tcell.KeyCtrlN:
		dest_box.SetText("NNN")
	case tcell.KeyCtrlP:
		dest_box.SetText("PPP")
	case tcell.KeyCtrlT:
		dest_box.SetText("TTT")
	case tcell.KeyCtrlS:
		dest_box.SetText("SSS")
	}

	return event
}
