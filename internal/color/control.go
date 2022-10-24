package color

import (
	"github.com/gdamore/tcell/v2"
)

type windowStyle struct {
	borderColor string
}

type Style struct {
	src             windowStyle
	dst             windowStyle
	backgroundColor string
	foregroundColor string
	selectedColor   string
	prefixColor     string
	labelColor      string
	pressColor      string
	highLightColor  string
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
		highLightColor:  "orange",
	}
}

func (s Style) BackgroundColor() tcell.Color {
	if s.Transparent {
		return tcell.ColorDefault
	}
	return themes[s.Theme][s.backgroundColor]
}

func (s Style) ForegroundColor() tcell.Color {
	return themes[s.Theme][s.foregroundColor]
}

func (s Style) SelectedColor() tcell.Color {
	return themes[s.Theme][s.selectedColor]
}

func (s Style) PrefixColor() tcell.Color {
	return themes[s.Theme][s.prefixColor]
}

func (s Style) LabelColor() tcell.Color {
	return themes[s.Theme][s.labelColor]
}

func (s Style) PressColor() tcell.Color {
	return themes[s.Theme][s.pressColor]
}

func (s Style) HighLightColor() tcell.Color {
	return themes[s.Theme][s.highLightColor]
}

func (s Style) SrcBorderColor() tcell.Color {
	return themes[s.Theme][s.src.borderColor]
}

func (s Style) DstBorderColor() tcell.Color {
	return themes[s.Theme][s.dst.borderColor]
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
