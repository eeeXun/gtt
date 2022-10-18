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

	dest_box.SetBorder(true).
		SetTitle(translator.dest_lang).
		SetBorderColor(window.dest.border_color).
		SetTitleColor(window.dest.border_color).
		SetBackgroundColor(window.dest.background_color)
}
