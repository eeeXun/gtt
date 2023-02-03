package translate

import (
	"gtt/internal/lock"
	"gtt/internal/translate/argostranslate"
	"gtt/internal/translate/googletranslate"
)

var (
	AllTranslator = []string{"ArgosTranslate", "GoogleTranslate"}
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
	Translate(message string) (
		translation string,
		definition string,
		partOfSpeech string,
		err error)
	// text to speech
	LockAvailable() bool
	LockAcquire()
	StopTTS()
	PlayTTS(lang string, message string) error
}

func NewTranslator(name string) Translator {
	var translator Translator

	switch name {
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
