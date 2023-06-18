package ui

import (
	"github.com/gdamore/tcell/v2"
)

type keyData struct {
	ch  rune
	key tcell.Key
}

type KeyMaps map[string]keyData

func NewKeyData(chStr string) keyData {
	var key tcell.Key
	ch := rune(chStr[0])

	switch ch {
	case ' ':
		key = tcell.KeyCtrlSpace
	case '\\':
		key = tcell.KeyCtrlBackslash
	case ']':
		key = tcell.KeyCtrlRightSq
	case '^':
		key = tcell.KeyCtrlCarat
	case '_':
		key = tcell.KeyCtrlUnderscore
	default:
		// This should be a to z
		key = tcell.KeyCtrlA + tcell.Key(ch-'a')
	}

	return keyData{
		ch:  ch,
		key: key,
	}
}

func (k keyData) GetName() rune {
	return k.ch
}

func (k keyData) GetKey() tcell.Key {
	return k.key
}
