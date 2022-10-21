package main

import (
	"github.com/gdamore/tcell/v2"
)

var (
	themes_name             = []string{"Gruvbox", "Nord"}
	Transparent tcell.Color = tcell.ColorDefault
	Themes                  = map[string]map[string]tcell.Color{
		"Gruvbox": {
			"bg":     tcell.NewHexColor(0x282828),
			"fg":     tcell.NewHexColor(0xebdbb2),
			"gray":   tcell.NewHexColor(0x665c54),
			"red":    tcell.NewHexColor(0xfb4934),
			"green":  tcell.NewHexColor(0xfabd2f),
			"yellow": tcell.NewHexColor(0xfabd2f),
			"blue":   tcell.NewHexColor(0x83a598),
			"purple": tcell.NewHexColor(0xd3869b),
			"cyan":   tcell.NewHexColor(0x8ec07c),
		},
		"Nord": {
			"bg":     tcell.NewHexColor(0x3b4252),
			"fg":     tcell.NewHexColor(0xeceff4),
			"gray":   tcell.NewHexColor(0x4c566a),
			"red":    tcell.NewHexColor(0xbf616a),
			"green":  tcell.NewHexColor(0xa3be8c),
			"yellow": tcell.NewHexColor(0xebcb8b),
			"blue":   tcell.NewHexColor(0x81a1c1),
			"purple": tcell.NewHexColor(0xb48ead),
			"cyan":   tcell.NewHexColor(0x8fbcbb),
		},
	}
)

type Colors struct {
	background_color tcell.Color
	foreground_color tcell.Color
	border_color     tcell.Color
	text_color       tcell.Color
	selected_color   tcell.Color
	prefix_color     tcell.Color
}

type Window struct {
	src Colors
	dst Colors
}

func (w *Window) color_init() {
	theme := "Gruvbox"
	transparent := true
	if transparent {
		w.src.background_color = Transparent
		w.dst.background_color = Transparent
	} else {
		w.src.background_color = Themes[theme]["bg"]
		w.dst.background_color = Themes[theme]["bg"]
	}
	w.src.border_color = Themes[theme]["red"]
	w.src.foreground_color = Themes[theme]["fg"]
	w.src.selected_color = Themes[theme]["gray"]
	w.src.prefix_color = Themes[theme]["yellow"]
	w.dst.foreground_color = Themes[theme]["fg"]
	w.dst.border_color = Themes[theme]["blue"]
	w.dst.prefix_color = Themes[theme]["yellow"]
}
