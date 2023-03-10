package translate

import (
	"github.com/eeeXun/gtt/internal/translate/apertiumtranslate"
	"github.com/eeeXun/gtt/internal/translate/argostranslate"
	"github.com/eeeXun/gtt/internal/translate/googletranslate"
	"github.com/eeeXun/gtt/internal/translate/reversotranslate"
)

var (
	AllTranslator = []string{"ApertiumTranslate", "ArgosTranslate", "GoogleTranslate", "ReversoTranslate"}
)

type Translator interface {
	// Get engine name of the translator
	GetEngineName() string

	// Get all languages of the translator
	GetAllLang() []string

	// Get source language of the translator
	GetSrcLang() string

	// Get destination language of the translator
	GetDstLang() string

	// Set source language of the translator
	SetSrcLang(lang string)

	// Set destination language of the translator
	SetDstLang(lang string)

	// Swap source and destination language of the translator
	SwapLang()

	// Check if lock is available
	LockAvailable() bool

	// Acquire the lock
	AcquireLock()

	// Stop text to speech
	StopTTS()

	// Translate from source to destination language
	Translate(message string) (translation, definition, partOfSpeech string, err error)

	// Play text to speech
	PlayTTS(lang, message string) error
}

func NewTranslator(name string) Translator {
	var translator Translator

	switch name {
	case "ApertiumTranslate":
		translator = apertiumtranslate.NewApertiumTranslate()
	case "ArgosTranslate":
		translator = argostranslate.NewArgosTranslate()
	case "GoogleTranslate":
		translator = googletranslate.NewGoogleTranslate()
	case "ReversoTranslate":
		translator = reversotranslate.NewReversoTranslate()
	}

	return translator
}
