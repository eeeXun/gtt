package google

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

const (
	ttsURL = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

func (t *GoogleTranslate) LockAvailable() bool {
	return t.SoundLock.Available()
}

func (t *GoogleTranslate) LockAcquire() {
	t.SoundLock.Acquire()
}

func (t *GoogleTranslate) StopTTS() {
	t.SoundLock.Stop = true
}

func (t *GoogleTranslate) PlayTTS(lang string, message string) error {
	urlStr := fmt.Sprintf(
		ttsURL,
		url.QueryEscape(message),
		langCode[lang],
	)
	res, err := http.Get(urlStr)
	if err != nil {
		t.SoundLock.Release()
		return err
	}
	decoder, err := mp3.NewDecoder(res.Body)
	if err != nil {
		t.SoundLock.Release()
		return err
	}
	otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		t.SoundLock.Release()
		return err
	}
	<-readyChan
	player := otoCtx.NewPlayer(decoder)
	player.Play()
	for player.IsPlaying() {
		if t.SoundLock.Stop {
			t.SoundLock.Release()
			return nil
		} else {
			time.Sleep(time.Millisecond)
		}
	}
	if err = player.Close(); err != nil {
		t.SoundLock.Release()
		return err
	}

	t.SoundLock.Release()
	return nil
}
