package argostranslate

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
	textURL = "https://translate.argosopentech.com/translate"
)

type ArgosTranslate struct {
	srcLang    string
	dstLang    string
	EngineName string
	SoundLock  *lock.Lock
}

func (t *ArgosTranslate) GetEngineName() string {
	return t.EngineName
}

func (t *ArgosTranslate) GetAllLang() []string {
	return lang
}

func (t *ArgosTranslate) GetSrcLang() string {
	return t.srcLang
}

func (t *ArgosTranslate) GetDstLang() string {
	return t.dstLang
}

func (t *ArgosTranslate) SetSrcLang(srcLang string) {
	t.srcLang = srcLang
}

func (t *ArgosTranslate) SetDstLang(dstLang string) {
	t.dstLang = dstLang
}

func (t *ArgosTranslate) SwapLang() {
	t.srcLang, t.dstLang = t.dstLang, t.srcLang
}

func (t *ArgosTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data map[string]interface{}

	res, err := http.PostForm(textURL,
		url.Values{
			"q":      {message},
			"source": {langCode[t.srcLang]},
			"target": {langCode[t.dstLang]},
		})
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
		translation += fmt.Sprintf("%v", data["translatedText"])

		return translation, definition, partOfSpeech, nil
	}
	return "", "", "", errors.New("Translation not found")
}
