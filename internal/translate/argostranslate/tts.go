package argostranslate

import (
	"errors"
)

func (t *ArgosTranslate) LockAvailable() bool {
	return t.SoundLock.Available()
}

func (t *ArgosTranslate) LockAcquire() {
	t.SoundLock.Acquire()
}

func (t *ArgosTranslate) StopTTS() {
	t.SoundLock.Stop = true
}

func (t *ArgosTranslate) PlayTTS(lang string, message string) error {
	t.SoundLock.Release()
	return errors.New(t.EngineName + " does not support text to speech")
}
