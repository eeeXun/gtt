package apertiumtranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/eeeXun/gtt/internal/lock"
)

const (
	textURL = "https://www.apertium.org/apy/translate?langpair=%s|%s&q=%s"
)

type ApertiumTranslate struct {
	srcLang    string
	dstLang    string
	EngineName string
	SoundLock  *lock.Lock
}

func (t *ApertiumTranslate) GetEngineName() string {
	return t.EngineName
}

func (t *ApertiumTranslate) GetAllLang() []string {
	return lang
}

func (t *ApertiumTranslate) GetSrcLang() string {
	return t.srcLang
}

func (t *ApertiumTranslate) GetDstLang() string {
	return t.dstLang
}

func (t *ApertiumTranslate) SetSrcLang(srcLang string) {
	t.srcLang = srcLang
}

func (t *ApertiumTranslate) SetDstLang(dstLang string) {
	t.dstLang = dstLang
}

func (t *ApertiumTranslate) SwapLang() {
	t.srcLang, t.dstLang = t.dstLang, t.srcLang
}

func (t *ApertiumTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data map[string]interface{}

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
		switch res.StatusCode {
		case 200:
			translation += fmt.Sprintf("%v",
				data["responseData"].(map[string]interface{})["translatedText"])
		default:
			return "", "", "", errors.New(
				fmt.Sprintf("%s does not support translate from %s to %s.\nSee available pair on %s",
					t.EngineName,
					t.srcLang,
					t.dstLang,
					"https://www.apertium.org/",
				))
		}

		return translation, definition, partOfSpeech, nil
	}
	return "", "", "", errors.New("Translation not found")
}
