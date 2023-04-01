package lingvatranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	// "time"

	"github.com/eeeXun/gtt/internal/translate/core"
	// "github.com/hajimehoshi/go-mp3"
	// "github.com/hajimehoshi/oto/v2"
)

const (
	textURL = "https://lingva.ml/_next/data/3qnDcUVykFKnSC3cdRX2t/%s/%s/%s.json"
)

type LingvaTranslate struct {
	*core.Language
	*core.TTSLock
	core.EngineName
}

func NewLingvaTranslate() *LingvaTranslate {
	return &LingvaTranslate{
		Language:   core.NewLanguage(),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("LingvaTranslate"),
	}
}

func (t *LingvaTranslate) GetAllLang() []string {
	return lang
}

func (t *LingvaTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data map[string]interface{}

	urlStr := fmt.Sprintf(
		textURL,
		langCode[t.GetSrcLang()],
		langCode[t.GetDstLang()],
		url.PathEscape(message),
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

	data = data["pageProps"].(map[string]interface{})
	// translation
	translation = fmt.Sprintf("%v", data["translation"])
	// definition
	for _, definitions := range data["info"].(map[string]interface{})["definitions"].([]interface{}) {
		definitions := definitions.(map[string]interface{})
		// part of speech
		if pos, ok := definitions["type"]; ok {
			definition += fmt.Sprintf("[%v]\n", pos)
		} else {
			definition += "[]\n"
		}
		for _, sentences := range definitions["list"].([]interface{}) {
			sentences := sentences.(map[string]interface{})
			// definition
			def := sentences["definition"]
			definition += fmt.Sprintf("\t- %v\n", def)
			// example sentence
			if example, ok := sentences["example"]; ok {
				definition += fmt.Sprintf("\t\t\"%v\"\n", example)
			}
		}
	}
	// part of speech
	for _, partOfSpeeches := range data["info"].(map[string]interface{})["extraTranslations"].([]interface{}) {
		partOfSpeeches := partOfSpeeches.(map[string]interface{})
		// part of speech
		if pos, ok := partOfSpeeches["type"]; ok {
			partOfSpeech += fmt.Sprintf("[%v]\n", pos)
		} else {
			partOfSpeech += "[]\n"
		}
		for _, words := range partOfSpeeches["list"].([]interface{}) {
			words := words.(map[string]interface{})
			dstWord := words["word"]
			partOfSpeech += fmt.Sprintf("\t%v:", dstWord)
			// src lang
			firstWord := true
			for _, word := range words["meanings"].([]interface{}) {
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

	return translation, definition, partOfSpeech, nil
}

func (t *LingvaTranslate) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return nil
}
