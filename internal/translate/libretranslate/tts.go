package libretranslate

import (
	"errors"
)

const (
	ttsURL = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

func (t *LibreTranslate) LockAvailable() bool {
	return t.SoundLock.Available()
}

func (t *LibreTranslate) LockAcquire() {
	t.SoundLock.Acquire()
}

func (t *LibreTranslate) StopTTS() {
	t.SoundLock.Stop = true
}

func (t *LibreTranslate) PlayTTS(lang string, message string) error {
	t.SoundLock.Release()
	return errors.New("LibreTranslate does not support text to speech")
}
