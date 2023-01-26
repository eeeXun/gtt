package googletranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"gtt/internal/lock"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	textURL = "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t&dt=bd&dt=md&dt=ex&sl=%s&tl=%s&q=%s"
)

type GoogleTranslate struct {
	srcLang   string
	dstLang   string
	SoundLock *lock.Lock
}

func (t *GoogleTranslate) GetAllLang() []string {
	return lang
}

func (t *GoogleTranslate) GetSrcLang() string {
	return t.srcLang
}

func (t *GoogleTranslate) GetDstLang() string {
	return t.dstLang
}

func (t *GoogleTranslate) SetSrcLang(srcLang string) {
	t.srcLang = srcLang
}

func (t *GoogleTranslate) SetDstLang(dstLang string) {
	t.dstLang = dstLang
}

func (t *GoogleTranslate) SwapLang() {
	t.srcLang, t.dstLang = t.dstLang, t.srcLang
}

func (t *GoogleTranslate) Translate(message string) (
	translation string,
	definition string,
	partOfSpeech string,
	err error) {
	var data []interface{}

	urlStr := fmt.Sprintf(
		textURL,
		langCode[t.srcLang],
		langCode[t.dstLang],
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

	if len(data) > 0 {
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

	return "", "", "", errors.New("Translation not found")
}
