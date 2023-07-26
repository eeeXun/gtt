package style

import (
	"github.com/gdamore/tcell/v2"
)

var (
	AllTheme = []string{"gruvbox", "nord"}
	Palette  = []string{"red", "green", "yellow", "blue", "purple", "cyan", "orange"}
	themes   = map[string]map[string]tcell.Color{
		"gruvbox": {
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
		"nord": {
			"bg":     tcell.NewHexColor(0x2e3440),
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

func NewTheme(name string, palette map[string]int32) {
	AllTheme = append(AllTheme, name)
	themes[name] = make(map[string]tcell.Color)
	for color, rgb := range palette {
		themes[name][color] = tcell.NewHexColor(rgb)
	}
}
