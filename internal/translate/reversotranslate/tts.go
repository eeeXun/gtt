package reversotranslate

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

const (
	ttsURL = "https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName=%s?voiceSpeed=80&inputText=%s"
)

func (t *ReversoTranslate) LockAvailable() bool {
	return t.SoundLock.Available()
}

func (t *ReversoTranslate) LockAcquire() {
	t.SoundLock.Acquire()
}

func (t *ReversoTranslate) StopTTS() {
	t.SoundLock.Stop = true
}

func (t *ReversoTranslate) PlayTTS(lang, message string) error {
	name, ok := voiceName[lang]
	if !ok {
		return errors.New(t.EngineName + " does not support text to speech of " + lang)
	}
	urlStr := fmt.Sprintf(
		ttsURL,
		name,
		base64.StdEncoding.EncodeToString([]byte(message)),
	)
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Add("User-Agent", userAgent)
	res, err := http.DefaultClient.Do(req)
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
