package main

import (
	"github.com/gdamore/tcell/v2"
)

var (
	themesName             = []string{"Gruvbox", "Nord"}
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
	backgroundColor tcell.Color
	foregroundColor tcell.Color
	borderColor     tcell.Color
	textColor       tcell.Color
	selectedColor   tcell.Color
	prefixColor     tcell.Color
	labelColor      tcell.Color
	pressColor      tcell.Color
}

type Window struct {
	src Colors
	dst Colors
}

func (w *Window) colorInit() {
	if transparent {
		w.src.backgroundColor = Transparent
		w.dst.backgroundColor = Transparent
	} else {
		w.src.backgroundColor = Themes[theme]["bg"]
		w.dst.backgroundColor = Themes[theme]["bg"]
	}
	w.src.borderColor = Themes[theme]["red"]
	w.src.foregroundColor = Themes[theme]["fg"]
	w.src.selectedColor = Themes[theme]["gray"]
	w.src.prefixColor = Themes[theme]["cyan"]
	w.src.pressColor = Themes[theme]["purple"]
	w.src.labelColor = Themes[theme]["yellow"]
	w.dst.foregroundColor = Themes[theme]["fg"]
	w.dst.selectedColor = Themes[theme]["gray"]
	w.dst.borderColor = Themes[theme]["blue"]
	w.dst.prefixColor = Themes[theme]["cyan"]
}
