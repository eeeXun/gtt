package translate

import (
	"github.com/eeeXun/gtt/internal/lock"
	"github.com/eeeXun/gtt/internal/translate/apertiumtranslate"
	"github.com/eeeXun/gtt/internal/translate/argostranslate"
	"github.com/eeeXun/gtt/internal/translate/googletranslate"
)

var (
	AllTranslator = []string{"ApertiumTranslate", "ArgosTranslate", "GoogleTranslate"}
)

type Translator interface {
	// engine name
	GetEngineName() string
	// text
	GetAllLang() []string
	GetSrcLang() string
	GetDstLang() string
	SetSrcLang(srcLang string)
	SetDstLang(dstLang string)
	SwapLang()
	Translate(message string) (translation, definition, partOfSpeech string, err error)
	// text to speech
	LockAvailable() bool
	LockAcquire()
	StopTTS()
	PlayTTS(lang, message string) error
}

func NewTranslator(name string) Translator {
	var translator Translator

	switch name {
	case "ApertiumTranslate":
		translator = &apertiumtranslate.ApertiumTranslate{
			EngineName: "ApertiumTranslate",
			SoundLock:  lock.NewLock(),
		}
	case "ArgosTranslate":
		translator = &argostranslate.ArgosTranslate{
			EngineName: "ArgosTranslate",
			SoundLock:  lock.NewLock(),
		}
	case "GoogleTranslate":
		translator = &googletranslate.GoogleTranslate{
			EngineName: "GoogleTranslate",
			SoundLock:  lock.NewLock(),
		}
	}

	return translator
}
