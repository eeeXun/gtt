package main

import (
	"github.com/gdamore/tcell/v2"
)

var keyNames = map[tcell.Key]string{
	tcell.KeyF1:             "F1",
	tcell.KeyF2:             "F2",
	tcell.KeyF3:             "F3",
	tcell.KeyF4:             "F4",
	tcell.KeyF5:             "F5",
	tcell.KeyF6:             "F6",
	tcell.KeyF7:             "F7",
	tcell.KeyF8:             "F8",
	tcell.KeyF9:             "F9",
	tcell.KeyF10:            "F10",
	tcell.KeyF11:            "F11",
	tcell.KeyF12:            "F12",
	tcell.KeyF13:            "F13",
	tcell.KeyF14:            "F14",
	tcell.KeyF15:            "F15",
	tcell.KeyF16:            "F16",
	tcell.KeyF17:            "F17",
	tcell.KeyF18:            "F18",
	tcell.KeyF19:            "F19",
	tcell.KeyF20:            "F20",
	tcell.KeyF21:            "F21",
	tcell.KeyF22:            "F22",
	tcell.KeyF23:            "F23",
	tcell.KeyF24:            "F24",
	tcell.KeyF25:            "F25",
	tcell.KeyF26:            "F26",
	tcell.KeyF27:            "F27",
	tcell.KeyF28:            "F28",
	tcell.KeyF29:            "F29",
	tcell.KeyF30:            "F30",
	tcell.KeyF31:            "F31",
	tcell.KeyF32:            "F32",
	tcell.KeyF33:            "F33",
	tcell.KeyF34:            "F34",
	tcell.KeyF35:            "F35",
	tcell.KeyF36:            "F36",
	tcell.KeyF37:            "F37",
	tcell.KeyF38:            "F38",
	tcell.KeyF39:            "F39",
	tcell.KeyF40:            "F40",
	tcell.KeyF41:            "F41",
	tcell.KeyF42:            "F42",
	tcell.KeyF43:            "F43",
	tcell.KeyF44:            "F44",
	tcell.KeyF45:            "F45",
	tcell.KeyF46:            "F46",
	tcell.KeyF47:            "F47",
	tcell.KeyF48:            "F48",
	tcell.KeyF49:            "F49",
	tcell.KeyF50:            "F50",
	tcell.KeyF51:            "F51",
	tcell.KeyF52:            "F52",
	tcell.KeyF53:            "F53",
	tcell.KeyF54:            "F54",
	tcell.KeyF55:            "F55",
	tcell.KeyF56:            "F56",
	tcell.KeyF57:            "F57",
	tcell.KeyF58:            "F58",
	tcell.KeyF59:            "F59",
	tcell.KeyF60:            "F60",
	tcell.KeyF61:            "F61",
	tcell.KeyF62:            "F62",
	tcell.KeyF63:            "F63",
	tcell.KeyF64:            "F64",
	tcell.KeyCtrlA:          "C-a",
	tcell.KeyCtrlB:          "C-b",
	tcell.KeyCtrlC:          "C-c",
	tcell.KeyCtrlD:          "C-d",
	tcell.KeyCtrlE:          "C-e",
	tcell.KeyCtrlF:          "C-f",
	tcell.KeyCtrlG:          "C-g",
	tcell.KeyCtrlJ:          "C-j",
	tcell.KeyCtrlK:          "C-k",
	tcell.KeyCtrlL:          "C-l",
	tcell.KeyCtrlN:          "C-n",
	tcell.KeyCtrlO:          "C-o",
	tcell.KeyCtrlP:          "C-p",
	tcell.KeyCtrlQ:          "C-q",
	tcell.KeyCtrlR:          "C-r",
	tcell.KeyCtrlS:          "C-s",
	tcell.KeyCtrlT:          "C-t",
	tcell.KeyCtrlU:          "C-u",
	tcell.KeyCtrlV:          "C-v",
	tcell.KeyCtrlW:          "C-w",
	tcell.KeyCtrlX:          "C-x",
	tcell.KeyCtrlY:          "C-y",
	tcell.KeyCtrlZ:          "C-z",
	tcell.KeyCtrlSpace:      "C-Space",
	tcell.KeyCtrlUnderscore: "C-_",
	tcell.KeyCtrlRightSq:    "C-]",
	tcell.KeyCtrlBackslash:  "C-\\",
	tcell.KeyCtrlCarat:      "C-^",
}

func getKeyName(event *tcell.EventKey) string {
	var key = event.Key()

	keyName := keyNames[key]

	if event.Modifiers() == tcell.ModAlt && key == tcell.KeyRune {
		if event.Rune() == ' ' {
			keyName = "A-Space"
		} else {
			keyName = "A-" + string(event.Rune())
		}
	}

	return keyName
}
