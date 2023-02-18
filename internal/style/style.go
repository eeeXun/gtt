package style

import (
	"github.com/gdamore/tcell/v2"
)

type style struct {
	HideBelow       bool
	Transparent     bool
	Theme           string
	srcBorderColor  string
	dstBorderColor  string
	backgroundColor string
	foregroundColor string
	selectedColor   string
	prefixColor     string
	labelColor      string
	pressColor      string
	highLightColor  string
}

func NewStyle() *style {
	return &style{
		backgroundColor: "bg",
		foregroundColor: "fg",
		selectedColor:   "gray",
		prefixColor:     "cyan",
		labelColor:      "yellow",
		pressColor:      "purple",
		highLightColor:  "orange",
	}
}

func (s style) BackgroundColor() tcell.Color {
	if s.Transparent {
		return tcell.ColorDefault
	}
	return themes[s.Theme][s.backgroundColor]
}

func (s style) ForegroundColor() tcell.Color {
	return themes[s.Theme][s.foregroundColor]
}

func (s style) SelectedColor() tcell.Color {
	return themes[s.Theme][s.selectedColor]
}

func (s style) PrefixColor() tcell.Color {
	return themes[s.Theme][s.prefixColor]
}

func (s style) LabelColor() tcell.Color {
	return themes[s.Theme][s.labelColor]
}

func (s style) PressColor() tcell.Color {
	return themes[s.Theme][s.pressColor]
}

func (s style) HighLightColor() tcell.Color {
	return themes[s.Theme][s.highLightColor]
}

func (s style) SrcBorderColor() tcell.Color {
	return themes[s.Theme][s.srcBorderColor]
}

func (s style) DstBorderColor() tcell.Color {
	return themes[s.Theme][s.dstBorderColor]
}

func (s style) SrcBorderStr() string {
	return s.srcBorderColor
}

func (s style) DstBorderStr() string {
	return s.dstBorderColor
}

func (s *style) SetSrcBorderColor(color string) *style {
	s.srcBorderColor = color
	return s
}

func (s *style) SetDstBorderColor(color string) *style {
	s.dstBorderColor = color
	return s
}
