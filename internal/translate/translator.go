package translate

import (
	"gtt/internal/lock"
	"gtt/internal/translate/google"
)

type Translator interface {
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

func NewGoogleTranslate() *google.GoogleTranslate {
	return &google.GoogleTranslate{
		SoundLock: lock.NewLock(),
	}
}
