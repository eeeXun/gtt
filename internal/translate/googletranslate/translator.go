package googletranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/eeeXun/gtt/internal/translate/core"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

const (
	textURL = "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t&dt=bd&dt=md&dt=ex&sl=%s&tl=%s&q=%s"
	ttsURL  = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

type GoogleTranslate struct {
	*core.Language
	*core.TTSLock
	core.EngineName
}

func NewGoogleTranslate() *GoogleTranslate {
	return &GoogleTranslate{
		Language:   core.NewLanguage(),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("GoogleTranslate"),
	}
}

func (t *GoogleTranslate) GetAllLang() []string {
	return lang
}

func (t *GoogleTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data []interface{}

	urlStr := fmt.Sprintf(
		textURL,
		langCode[t.GetSrcLang()],
		langCode[t.GetDstLang()],
		url.QueryEscape(message),
	)
	res, err := http.Get(urlStr)
	if err != nil {
		return "", "", "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", "", err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", "", "", err
	}

	if len(data) <= 0 {
		return "", "", "", errors.New("Translation not found")
	}

	// translation = data[0]
	for _, lines := range data[0].([]interface{}) {
		translatedLine := lines.([]interface{})[0]
		translation += fmt.Sprintf("%v", translatedLine)
	}
	// part of speech = data[1]
	if data[1] != nil {
		for _, parts := range data[1].([]interface{}) {
			// part of speech
			part := parts.([]interface{})[0]
			partOfSpeech += fmt.Sprintf("[%v]\n", part)
			for _, words := range parts.([]interface{})[2].([]interface{}) {
				// dst lang
				dstWord := words.([]interface{})[0]
				partOfSpeech += fmt.Sprintf("\t%v:", dstWord)
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
		}
	}
	// definition = data[12]
	if len(data) >= 13 && data[12] != nil {
		for _, parts := range data[12].([]interface{}) {
			// part of speech
			part := parts.([]interface{})[0]
			definition += fmt.Sprintf("[%v]\n", part)
			for _, sentences := range parts.([]interface{})[1].([]interface{}) {
				// definition
				def := sentences.([]interface{})[0]
				definition += fmt.Sprintf("\t- %v\n", def)
				// example sentence
				if len(sentences.([]interface{})) >= 3 && sentences.([]interface{})[2] != nil {
					example := sentences.([]interface{})[2]
					definition += fmt.Sprintf("\t\t\"%v\"\n", example)
				}
			}
		}
	}

	return translation, definition, partOfSpeech, nil
}

func (t *GoogleTranslate) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	urlStr := fmt.Sprintf(
		ttsURL,
		url.QueryEscape(message),
		langCode[lang],
	)
	res, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	decoder, err := mp3.NewDecoder(res.Body)
	if err != nil {
		return err
	}
	otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-readyChan
	player := otoCtx.NewPlayer(decoder)
	player.Play()
	for player.IsPlaying() {
		if t.IsStopped() {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	if err = player.Close(); err != nil {
		return err
	}

	return nil
}
