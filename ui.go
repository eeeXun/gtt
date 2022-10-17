package main

import (
	"github.com/gdamore/tcell/v2"
	// "github.com/rivo/tview"
)

type Colors struct {
	background_color tcell.Color
	border_color     tcell.Color
	foreground_color tcell.Color
	text_color       tcell.Color
}

type Window struct {
	src  Colors
	dest Colors
}

func (w *Window) color_init() {
	w.src.background_color = tcell.ColorDefault
	w.src.border_color = tcell.NewHexColor(0xFBF1C7)
	w.src.foreground_color = tcell.NewHexColor(0xFBF1C7)
	w.dest.background_color = tcell.ColorDefault
	w.dest.border_color = tcell.NewHexColor(0xFBF1C7)
	w.dest.foreground_color = tcell.NewHexColor(0xFBF1C7)
}

func ui_init() {
	src_box.SetBorder(true).
		SetTitle(translator.src_lang).
		SetBorderColor(window.src.border_color).
		SetTitleColor(window.src.border_color)
	src_box.SetBackgroundColor(window.src.background_color)
	src_box.SetTextStyle(tcell.StyleDefault.
		Background(window.src.background_color).
		Foreground(window.src.foreground_color))

	dest_box.SetBorder(true).
		SetTitle(translator.dest_lang).
		SetBorderColor(window.dest.border_color).
		SetTitleColor(window.dest.border_color)
	dest_box.SetBackgroundColor(window.dest.background_color)
}
