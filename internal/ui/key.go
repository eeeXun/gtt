package ui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

type keyData struct {
	name string
	key  tcell.Key
}

type KeyMaps map[string]keyData

func NewKeyData(keyStr string) keyData {
	var (
		name string
		key  tcell.Key
	)

	if len(keyStr) > 1 && keyStr[0] == 'F' {
		// function key, can be F1 to F64
		name = keyStr
		fNum, err := strconv.Atoi(keyStr[1:])
		if err != nil {
			panic(err)
		}
		key = tcell.KeyF1 + tcell.Key(fNum-1)
	} else {
		switch keyStr[0] {
		case ' ':
			name = "C-Space"
			key = tcell.KeyCtrlSpace
		case '\\':
			name = "C-\\"
			key = tcell.KeyCtrlBackslash
		case ']':
			name = "C-]"
			key = tcell.KeyCtrlRightSq
		case '^':
			name = "C-^"
			key = tcell.KeyCtrlCarat
		case '_':
			name = "C-_"
			key = tcell.KeyCtrlUnderscore
		default:
			// This should be a to z
			name = "C-" + keyStr
			key = tcell.KeyCtrlA + tcell.Key(keyStr[0]-'a')
		}
	}

	return keyData{
		name: name,
		key:  key,
	}
}

func (k keyData) GetName() string {
	return k.name
}

func (k keyData) GetKey() tcell.Key {
	return k.key
}
