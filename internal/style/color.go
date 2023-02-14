package style

import (
	"github.com/gdamore/tcell/v2"
)

var (
	AllTheme = []string{"Gruvbox", "Nord"}
	Palette  = []string{"red", "green", "yellow", "blue", "purple", "cyan", "orange"}
	themes   = map[string]map[string]tcell.Color{
		"Gruvbox": {
			"bg":     tcell.NewHexColor(0x282828),
			"fg":     tcell.NewHexColor(0xebdbb2),
			"gray":   tcell.NewHexColor(0x665c54),
			"red":    tcell.NewHexColor(0xfb4934),
			"green":  tcell.NewHexColor(0xb8bb26),
			"yellow": tcell.NewHexColor(0xfabd2f),
			"blue":   tcell.NewHexColor(0x83a598),
			"purple": tcell.NewHexColor(0xd3869b),
			"cyan":   tcell.NewHexColor(0x8ec07c),
			"orange": tcell.NewHexColor(0xfe8019),
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
			"orange": tcell.NewHexColor(0xd08770),
		},
	}
)
