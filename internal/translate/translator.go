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

func NewArgosTranslate() *argostranslate.ArgosTranslate {
	return &argostranslate.ArgosTranslate{
		EngineName: "ArgosTranslate",
		SoundLock:  lock.NewLock(),
	}
}

func NewGoogleTranslate() *googletranslate.GoogleTranslate {
	return &googletranslate.GoogleTranslate{
		EngineName: "GoogleTranslate",
		SoundLock:  lock.NewLock(),
	}
}
