package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	textURL  = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s"
	soundURL = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

type Lock struct {
	stop        bool
	threadCount int8
}

func NewLock() *Lock {
	return &Lock{
		stop:        true,
		threadCount: 0,
	}
}

func (l *Lock) Available() bool {
	return l.stop && l.threadCount == 0
}

func (l *Lock) Acquire() {
	l.stop = false
	l.threadCount++
}

func (l *Lock) Release() {
	l.stop = true
	l.threadCount--
}

type Translator struct {
	srcLang   string
	dstLang   string
	soundLock *Lock
}

func NewTranslator() *Translator {
	return &Translator{
		soundLock: NewLock(),
	}
}

func (t *Translator) Translate(message string) (string, error) {
	var data []interface{}
	var translated string

	url_str := fmt.Sprintf(
		textURL,
		LangCode[t.srcLang],
		LangCode[t.dstLang],
		url.QueryEscape(message),
	)
	res, err := http.Get(url_str)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if len(data) > 0 {
		result := data[0]
		for _, lines := range result.([]interface{}) {
			translated_line := lines.([]interface{})[0]
			translated += fmt.Sprintf("%v", translated_line)
		}
		return translated, nil
	}

	return "", errors.New("Translation not found")
}

func (t *Translator) PlaySound(lang string, message string) error {
	url_str := fmt.Sprintf(
		soundURL,
		url.QueryEscape(message),
		LangCode[lang],
	)
	res, err := http.Get(url_str)
	if err != nil {
		t.soundLock.Release()
		return err
	}
	decoder, err := mp3.NewDecoder(res.Body)
	if err != nil {
		t.soundLock.Release()
		return err
	}
	otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		t.soundLock.Release()
		return err
	}
	<-readyChan
	player := otoCtx.NewPlayer(decoder)
	player.Play()
	for player.IsPlaying() {
		if t.soundLock.stop {
			t.soundLock.Release()
			return nil
		} else {
			time.Sleep(time.Millisecond)
		}
	}
	if err = player.Close(); err != nil {
		t.soundLock.Release()
		return err
	}

	t.soundLock.Release()
	return nil
}
