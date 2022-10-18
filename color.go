package main

import (
	"github.com/gdamore/tcell/v2"
)

type palette struct {
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
	themes_name = []string{"Gruvbox", "Nord"}
	Themes      = map[string]palette{
		"Gruvbox": {
			bg:     tcell.NewHexColor(0x282828),
			fg:     tcell.NewHexColor(0xebdbb2),
			gray:   tcell.NewHexColor(0x928374),
			red:    tcell.NewHexColor(0xfb4934),
			green:  tcell.NewHexColor(0xfabd2f),
			yellow: tcell.NewHexColor(0xfabd2f),
			blue:   tcell.NewHexColor(0x83a598),
			purple: tcell.NewHexColor(0xd3869b),
			cyan:   tcell.NewHexColor(0x8ec07c),
		},
		"Nord": {
			bg:     tcell.NewHexColor(0x3b4252),
			fg:     tcell.NewHexColor(0xeceff4),
			gray:   tcell.NewHexColor(0x4c566a),
			red:    tcell.NewHexColor(0xbf616a),
			green:  tcell.NewHexColor(0xa3be8c),
			yellow: tcell.NewHexColor(0xebcb8b),
			blue:   tcell.NewHexColor(0x81a1c1),
			purple: tcell.NewHexColor(0xb48ead),
			cyan:   tcell.NewHexColor(0x8fbcbb),
		},
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
	w.src.background_color = Themes["Gruvbox"].bg
	w.src.border_color = Themes["Gruvbox"].red
	w.src.foreground_color = Themes["Gruvbox"].fg
	w.src.selected_color = Themes["Gruvbox"].gray
	w.dest.background_color = Themes["Gruvbox"].bg
	w.dest.border_color = Themes["Gruvbox"].blue
	w.dest.foreground_color = Themes["Gruvbox"].fg
}
