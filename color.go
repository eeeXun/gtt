package main

import (
	"github.com/gdamore/tcell/v2"
)

type color_theme struct {
	bg     tcell.Color
	fg     tcell.Color
	gray   tcell.Color
	red    tcell.Color
	green  tcell.Color
	yellow tcell.Color
	blue   tcell.Color
	purple tcell.Color
	cyan   tcell.Color
}

var (
	Gruvbox color_theme = color_theme{
		bg:     tcell.NewHexColor(0x282828),
		fg:     tcell.NewHexColor(0xebdbb2),
		gray:   tcell.NewHexColor(0x928374),
		red:    tcell.NewHexColor(0xfb4934),
		green:  tcell.NewHexColor(0xfabd2f),
		yellow: tcell.NewHexColor(0xfabd2f),
		blue:   tcell.NewHexColor(0x83a598),
		purple: tcell.NewHexColor(0xd3869b),
		cyan:   tcell.NewHexColor(0x8ec07c),
	}
	Nord color_theme = color_theme{
		bg:     tcell.NewHexColor(0x3b4252),
		fg:     tcell.NewHexColor(0xeceff4),
		gray:   tcell.NewHexColor(0x4c566a),
		red:    tcell.NewHexColor(0xbf616a),
		green:  tcell.NewHexColor(0xa3be8c),
		yellow: tcell.NewHexColor(0xebcb8b),
		blue:   tcell.NewHexColor(0x81a1c1),
		purple: tcell.NewHexColor(0xb48ead),
		cyan:   tcell.NewHexColor(0x8fbcbb),
	}
)

type Colors struct {
	background_color tcell.Color
	border_color     tcell.Color
	foreground_color tcell.Color
	text_color       tcell.Color
	selected_color   tcell.Color
}

type Window struct {
	src  Colors
	dest Colors
}

func (w *Window) color_init() {
	w.src.background_color = Gruvbox.bg
	w.src.border_color = Gruvbox.red
	w.src.foreground_color = Gruvbox.fg
	w.src.selected_color = Gruvbox.gray
	w.dest.background_color = Gruvbox.bg
	w.dest.border_color = Gruvbox.blue
	w.dest.foreground_color = Gruvbox.fg
}
