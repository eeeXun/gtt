package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

const (
	textURL  = "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t&dt=bd&sl=%s&tl=%s&q=%s"
	soundURL = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

type Translator struct {
	SrcLang   string
	DstLang   string
	SoundLock *Lock
}

func NewTranslator() *Translator {
	return &Translator{
		SoundLock: NewLock(),
	}
}

func (t *Translator) Translate(message string) (translation string, partOfSpeech string, err error) {
	var data []interface{}

	urlStr := fmt.Sprintf(
		textURL,
		LangCode[t.SrcLang],
		LangCode[t.DstLang],
		url.QueryEscape(message),
	)
	res, err := http.Get(urlStr)
	if err != nil {
		return "", "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return "", "", err
	}

	if len(data) > 0 {
		// translation
		for _, lines := range data[0].([]interface{}) {
			translatedLine := lines.([]interface{})[0]
			translation += fmt.Sprintf("%v", translatedLine)
		}

		// part of speech
		if data[1] != nil {
			for _, parts := range data[1].([]interface{}) {
				// part of speech
				partOfSpeech += fmt.Sprintf("[%v]\n", parts.([]interface{})[0])
				for _, words := range parts.([]interface{})[2].([]interface{}) {
					// dst lang
					partOfSpeech += fmt.Sprintf(
						"\t%v:", words.([]interface{})[0])
					// src lang
					firstWord := true
					for _, word := range words.([]interface{})[1].([]interface{}) {
						if firstWord {
							partOfSpeech += fmt.Sprintf(" %v", word)
							firstWord = false
						} else {
							partOfSpeech += fmt.Sprintf(", %v", word)
						}
					}
					partOfSpeech += "\n"
				}
				partOfSpeech += "\n"
			}
		}
		return translation, partOfSpeech, nil
	}

	return "", "", errors.New("Translation not found")
}

func (t *Translator) PlaySound(lang string, message string) error {
	urlStr := fmt.Sprintf(
		soundURL,
		url.QueryEscape(message),
		LangCode[lang],
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
