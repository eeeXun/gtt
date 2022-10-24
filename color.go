package main

import (
	"github.com/gdamore/tcell/v2"
)

var (
	AllTheme                = []string{"Gruvbox", "Nord"}
	Palette                 = []string{"red", "green", "yellow", "blue", "purple", "cyan"}
	Themes                  = map[string]map[string]tcell.Color{
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

type WindowStyle struct {
	borderColor string
}

type Style struct {
	src             WindowStyle
	dst             WindowStyle
	backgroundColor string
	foregroundColor string
	selectedColor   string
	prefixColor     string
	labelColor      string
	pressColor      string
	Theme           string
	Transparent     bool
}

func NewStyle() *Style {
	return &Style{
		backgroundColor: "bg",
		foregroundColor: "fg",
		selectedColor:   "gray",
		prefixColor:     "cyan",
		labelColor:      "yellow",
		pressColor:      "purple",
	}
}

func (s Style) BackgroundColor() tcell.Color {
	if s.Transparent {
		return tcell.ColorDefault
	}
	return Themes[s.Theme][s.backgroundColor]
}

func (s Style) ForegroundColor() tcell.Color {
	return Themes[s.Theme][s.foregroundColor]
}

func (s Style) SelectedColor() tcell.Color {
	return Themes[s.Theme][s.selectedColor]
}

func (s Style) PrefixColor() tcell.Color {
	return Themes[s.Theme][s.prefixColor]
}

func (s Style) LabelColor() tcell.Color {
	return Themes[s.Theme][s.labelColor]
}

func (s Style) PressColor() tcell.Color {
	return Themes[s.Theme][s.pressColor]
}

func (s Style) SrcBorderColor() tcell.Color {
	return Themes[s.Theme][s.src.borderColor]
}

func (s Style) DstBorderColor() tcell.Color {
	return Themes[s.Theme][s.dst.borderColor]
}

func (s Style) SrcBorderStr() string {
	return s.src.borderColor
}

func (s Style) DstBorderStr() string {
	return s.dst.borderColor
}

func (s *Style) SetSrcBorderColor(color string) *Style {
	s.src.borderColor = color
	return s
}

func (s *Style) SetDstBorderColor(color string) *Style {
	s.dst.borderColor = color
	return s
}
