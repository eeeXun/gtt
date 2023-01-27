package translate

import (
	"gtt/internal/lock"
	"gtt/internal/translate/googletranslate"
	"gtt/internal/translate/libretranslate"
)

var (
	AllTranslator = []string{"LibreTranslate", "GoogleTranslate"}
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

func NewGoogleTranslate() *googletranslate.GoogleTranslate {
	return &googletranslate.GoogleTranslate{
		EngineName: "GoogleTranslate",
		SoundLock:  lock.NewLock(),
	}
}

func NewLibreTranslate() *libretranslate.LibreTranslate {
	return &libretranslate.LibreTranslate{
		EngineName: "LibreTranslate",
		SoundLock:  lock.NewLock(),
	}
}
