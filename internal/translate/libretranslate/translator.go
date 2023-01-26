package libretranslate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gtt/internal/lock"
	"io/ioutil"
	"net/http"
)

const (
	textURL = "https://translate.argosopentech.com/translate"
)

type LibreTranslate struct {
	srcLang   string
	dstLang   string
	SoundLock *lock.Lock
}

func (t *LibreTranslate) GetAllLang() []string {
	return lang
}

func (t *LibreTranslate) GetSrcLang() string {
	return t.srcLang
}

func (t *LibreTranslate) GetDstLang() string {
	return t.dstLang
}

func (t *LibreTranslate) SetSrcLang(srcLang string) {
	t.srcLang = srcLang
}

func (t *LibreTranslate) SetDstLang(dstLang string) {
	t.dstLang = dstLang
}

func (t *LibreTranslate) SwapLang() {
	t.srcLang, t.dstLang = t.dstLang, t.srcLang
}

func (t *LibreTranslate) Translate(message string) (
	translation string,
	definition string,
	partOfSpeech string,
	err error) {
	var data interface{}

	res, err := http.Post(textURL,
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{
			"q": "%s",
			"source": "%s",
			"target": "%s" }`, message, langCode[t.srcLang], langCode[t.dstLang]))))
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

	if len(data.(map[string]interface{})) > 0 {
		translation += fmt.Sprintf("%v",
			data.(map[string]interface{})["translatedText"])

		return translation, definition, partOfSpeech, nil
	}
	return "", "", "", errors.New("Translation not found")
}
