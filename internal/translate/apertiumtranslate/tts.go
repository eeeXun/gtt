package apertiumtranslate

import (
	"errors"
)

func (t *ApertiumTranslate) LockAvailable() bool {
	return t.SoundLock.Available()
}

func (t *ApertiumTranslate) LockAcquire() {
	t.SoundLock.Acquire()
}

func (t *ApertiumTranslate) StopTTS() {
	t.SoundLock.Stop = true
}

func (t *ApertiumTranslate) PlayTTS(lang string, message string) error {
	t.SoundLock.Release()
	return errors.New(t.EngineName + " does not support text to speech")
}
